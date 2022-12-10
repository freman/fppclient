package fppclient

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func New(baseURL string, args ...newArg) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create new client tue to %w", err)
	}

	c := Client{
		baseURL: u,
	}

	for _, arg := range args {
		arg(&c)
	}

	// Define a default http client with a sane 10 second timeout.
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	return &c, nil
}

func constrainToByte(i int) int {
	if i > 255 {
		return 255
	}

	if i < 0 {
		return 0
	}

	return i
}

func (c Client) GetPlugins() (plugins Plugins, err error) {
	if err = c.httpGet("/api/plugin", &plugins); err != nil {
		return nil, fmt.Errorf("unable to retrieve plugins: %w", err)
	}

	return plugins, err
}

func (c Client) GetConfig(name string, v interface{}) error {
	if err := c.httpGet("/api/configfile/"+name, v); err != nil {
		return fmt.Errorf("unable to retrieve config file %q: %w", name, err)
	}

	return nil
}

func (c Client) GetChannelOutputs() (ChannelOutputs, error) {
	var resp ChannelOutputsObj
	if err := c.GetConfig("channeloutputs.json", &resp); err != nil {
		return nil, fmt.Errorf("unable to retrieve channel outputs: %w", err)
	}

	return resp.ChannelOutputs, nil
}

func (c Client) GetSchedule() (Schedule, error) {
	var resp ScheduleResponse
	if err := c.httpGet("/api/fppd/schedule", &resp); err != nil {
		return Schedule{}, fmt.Errorf("unable to retrieve schedule: %w", err)
	}

	return resp.Schedule, nil
}

type newArg func(c *Client)

func WithHTTPClient(httpClient *http.Client) newArg {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
