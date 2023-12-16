package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/manifoldco/promptui"

	"github.com/freman/fppclient"
)

func init() {
	// extend the default funcmap to add an add function
	promptui.FuncMap["add"] = func(a, b int) int {
		return a + b
	}
}

var cleanup []func()

func main() {
	signalHandler()
	defer shutdown()

	fppHost := flag.String("host", "10.0.0.249", "FPP host")

	flag.Parse()

	c, err := fppclient.New("http://" + *fppHost)
	if err != nil {
		panic(err)
	}

	model, err := promptForModel(c)
	if err != nil {
		panic(err)
	}

	outputs, err := c.GetChannelOutputs(context.TODO())
	if err != nil {
		panic(err)
	}

	outputPanel, err := filterOutputsByModel(model, outputs)
	if err != nil {
		panic(err)
	}

	fmt.Println("found", outputPanel.Type, "with", len(outputPanel.Panels), "panels")

	panel, err := promptForPanel(outputPanel.Panels)
	if err != nil {
		panic(err)
	}

	if err := c.SetOverlaysModelState(context.TODO(), model.Name, 1); err != nil {
		panic(err)
	}

	cleanup = append(cleanup, func() {
		if err := c.ClearOverlaysModel(context.TODO(), model.Name); err != nil {
			fmt.Println("Warning, failed to clear model:", err.Error())
		}

		// if you shut it down while it's clearing it leaves lit pixels
		time.Sleep(500 * time.Millisecond)

		if err := c.SetOverlaysModelState(context.TODO(), model.Name, 0); err != nil {
			fmt.Println("Warning, failed to turn off the panel:", err.Error())

		}
	})

	sequences := map[string][]int{
		"red":    {100, 0, 0},
		"green":  {0, 100, 0},
		"blue":   {0, 0, 100},
		"yellow": {80, 80, 0},
		"purple": {80, 0, 80},
		"cyan":   {0, 80, 80},
		"white":  {60, 60, 60},
	}

	type work struct {
		x, y  int
		color []int
	}

	chwork := make(chan work, 30)

	for i := 0; i < 9; i++ {
		go func() {
			for job := range chwork {
				if err := c.SetOverlaysModelPixel(context.TODO(), model.Name, job.x, job.y, job.color[0], job.color[1], job.color[2]); err != nil {
					fmt.Printf("Warning, failed to configure pixel %d,%d: %v", job.x, job.y, err)
				}
			}
		}()
	}

	for name, color := range sequences {
		fmt.Println("All", name)
		for x, xend := panel.XOffset, panel.XOffset+outputPanel.PanelWidth; x < xend; x++ {
			for y, yend := panel.YOffset, panel.YOffset+outputPanel.PanelHeight; y < yend; y++ {
				chwork <- work{
					x:     x,
					y:     y,
					color: color,
				}
			}
		}
		time.Sleep(10 * time.Second)
	}

	close(chwork)
}

func promptForModel(c *fppclient.Client) (model fppclient.Model, err error) {
	models, err := c.GetOverlaysModels(context.TODO())
	if err != nil {
		return model, err
	}

	idx, _, err := (&promptui.Select{
		Label: "Test which model?",
		Items: models,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Name }}",
			Inactive: "{{ .Name }}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .Name | faint }}`, promptui.IconGood),
			Active:   fmt.Sprintf("%s {{ .Name | underline }}", promptui.IconSelect),
			Details: `Channels: {{ .ChannelCount }}
StartChannel: {{ .StartChannel }}
Width: {{ .Width }}
Height: {{ .Height }}`,
		},
	}).Run()

	if err != nil {
		return model, err
	}

	return models[idx], nil
}

func filterOutputsByModel(model fppclient.Model, outputs fppclient.ChannelOutputs) (output fppclient.ChannelOutput, err error) {
	for _, o := range outputs {
		if o.StartChannel == model.StartChannel && o.ChannelCount == model.ChannelCount {
			return o, nil
		}
	}

	return output, fmt.Errorf("no matching output for model %s", model.Name)
}

func promptForPanel(panels fppclient.ChannelOutputPanels) (panel fppclient.ChannelOutputPanel, err error) {
	idx, _, err := (&promptui.Select{
		Label: "Test which panel?",
		Items: panels,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ add 1 .OutputNumber }}-{{ add 1 .PanelNumber }}",
			Inactive: "{{ add 1 .OutputNumber }}-{{ add 1 .PanelNumber }}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ add 1 .OutputNumber | faint }}{{ "-" | faint }}{{ add 1 .PanelNumber | faint }}`, promptui.IconGood),
			Active:   fmt.Sprintf(`%s {{  add 1 .OutputNumber | underline }}{{ "-" | underline }}{{ add 1 .PanelNumber | underline }}`, promptui.IconSelect),
		},
	}).Run()

	if err != nil {
		return panel, err
	}

	return panels[idx], nil
}

func shutdown() {
	fmt.Println("Shutting down.")
	for _, fn := range cleanup {
		fn()
	}
}

func signalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("^C intercepted.")
		shutdown()
		os.Exit(1)
	}()
}
