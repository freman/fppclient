package fppclient

import (
	"context"
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

func (c Client) GetPlugins(ctx context.Context) (plugins Plugins, err error) {
	if err = c.httpGet(ctx, "/api/plugin", &plugins); err != nil {
		return nil, fmt.Errorf("unable to retrieve plugins: %w", err)
	}

	return plugins, err
}

func (c Client) GetConfig(ctx context.Context, name string, v interface{}) error {
	if err := c.httpGet(ctx, "/api/configfile/"+name, v); err != nil {
		return fmt.Errorf("unable to retrieve config file %q: %w", name, err)
	}

	return nil
}

func (c Client) GetChannelOutputs(ctx context.Context) (ChannelOutputs, error) {
	var resp ChannelOutputsObj
	if err := c.GetConfig(ctx, "channeloutputs.json", &resp); err != nil {
		return nil, fmt.Errorf("unable to retrieve channel outputs: %w", err)
	}

	return resp.ChannelOutputs, nil
}

type newArg func(c *Client)

func WithHTTPClient(httpClient *http.Client) newArg {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
