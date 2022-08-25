package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/manifoldco/promptui"

	"github.com/freman/fppclient"
)

func main() {
	fppHost := flag.String("host", "10.0.0.249", "FPP host")

	flag.Parse()

	c, err := fppclient.New("http://" + *fppHost)
	if err != nil {
		panic(err)
	}

	models, err := c.GetModels()
	if err != nil {
		panic(err)
	}

	modelPrompt := promptui.Select{
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
	}

	idx, _, err := modelPrompt.Run()
	if err != nil {
		panic(err)
	}

	model := models[idx]
	nodes := model.ChannelCount / model.ChannelCountPerNode
	_ = nodes

	outputs, err := c.GetChannelOutputs()
	if err != nil {
		panic(err)
	}

	var outputPanel fppclient.ChannelOutput
	for _, output := range outputs {
		if output.StartChannel == model.StartChannel && output.ChannelCount == model.ChannelCount {
			outputPanel = output
			break
		}
	}

	if outputPanel.StartChannel == 0 {
		panic("panel config not found")
	}

	fmt.Println("found", outputPanel.Type, "with", len(outputPanel.Panels), "panels")

	panelPrompt := promptui.Select{
		Label: "Test which panel?",
		Items: outputPanel.Panels,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .OutputNumber }}-{{ .PanelNumber }}",
			Inactive: "{{ .OutputNumber }}-{{ .PanelNumber }}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .OutputNumber | faint }}{{ "-" | faint }}{{ .PanelNumber | faint }}`, promptui.IconGood),
			Active:   fmt.Sprintf(`%s {{  .OutputNumber | underline }}{{ "-" | underline }}{{ .PanelNumber | underline }}`, promptui.IconSelect),
		},
	}

	idx, _, err = panelPrompt.Run()
	if err != nil {
		panic(err)
	}

	panel := outputPanel.Panels[idx]
	startX := panel.XOffset
	startY := panel.YOffset
	endX := outputPanel.PanelWidth
	endY := outputPanel.PanelHeight

	c.SetModelState(model.Name, 1)

	sequences := map[string][]int{
		"red":    {100, 0, 0},
		"green":  {0, 100, 0},
		"blue":   {0, 0, 100},
		"yellow": {80, 80, 0},
		"purple": {80, 0, 80},
		"cyan":   {0, 80, 0},
		"white":  {60, 60, 60},
	}

	for name, color := range sequences {
		fmt.Println("All", name)
		for x := startX; x < endX; x++ {
			for y := startY; y < endY; y++ {
				c.SetModelPixel(model.Name, x, y, color[0], color[1], color[2])
			}
		}
		time.Sleep(10 * time.Second)
	}

	c.SetModelState(model.Name, 0)
}
