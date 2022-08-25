package fppclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

func (c Client) formatURL(path string) string {
	return c.baseURL.ResolveReference(
		&url.URL{
			Path: path,
		}).String()
}

func (c Client) httpGet(path string, v interface{}) error {
	u := c.formatURL(path)

	resp, err := c.httpClient.Get(u)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Help the server out by reading and discarding the body
		io.Copy(io.Discard, resp.Body)
		return fmt.Errorf("unexpected HTTP status %q (%d)", resp.Status, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("unable to parse response: %w", err)
	}

	return nil
}

func (c Client) httpPut(path string, in, out interface{}) error {
	u := c.formatURL(path)

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(in); err != nil {
		return fmt.Errorf("unable to marshal object: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, u, &buf)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Help the server out by reading and discarding the body
		io.Copy(io.Discard, resp.Body)
		return fmt.Errorf("unexpected HTTP status %q (%d)", resp.Status, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("unable to parse response: %w", err)
	}

	return nil
}

func (c Client) GetModels() (models Models, err error) {
	if err = c.httpGet("/api/overlays/models", &models); err != nil {
		return nil, fmt.Errorf("unable to retrieve models: %w", err)
	}

	return models, err
}

func (c Client) GetModel(name string) (*Model, error) {
	var model Model

	path := fmt.Sprintf("/api/overlays/model/%s", name)
	if err := c.httpGet(path, &model); err != nil {
		return nil, fmt.Errorf("unable to retrieve model %q: %w", name, err)
	}

	return &model, nil
}

func (c Client) ClearModel(name string) error {
	var resp Status

	path := fmt.Sprintf("/api/overlays/model/%s/clear", name)
	if err := c.httpGet(path, &resp); err != nil {
		return fmt.Errorf("unable to clear model %q: %w", name, err)
	}

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to clear model %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) GetModelData(name string, rle bool) (*ModelData, error) {
	var modelData ModelData

	path := fmt.Sprintf("/api/overlays/model/%s/data", name)
	if rle {
		path += "/rle"
	}

	if err := c.httpGet(path, &modelData); err != nil {
		return nil, fmt.Errorf("unable to retrieve model %q: %w", name, err)
	}

	return &modelData, nil
}

func (c Client) SetModelState(name string, state interface{}) error {
	path := fmt.Sprintf("/api/overlays/model/%s/state", name)
	//    $data = "{\"State\": ". $state . "}";

	var resp Status
	c.httpPut(path, struct {
		State interface{}
	}{State: state}, &resp)

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to clear model %q: %s", name, resp.Message)
	}

	return nil
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

func (c Client) FillModel(name string, r, g, b int) error {
	path := fmt.Sprintf("/api/overlays/model/%s/fill", name)

	fillReq := fillPixelRequest{
		RGB: []int{
			constrainToByte(r),
			constrainToByte(g),
			constrainToByte(b),
		},
	}

	var resp Status
	c.httpPut(path, &fillReq, &resp)

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to fill model %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) SetModelPixel(name string, x, y, r, g, b int) error {
	path := fmt.Sprintf("/api/overlays/model/%s/pixel", name)

	pixelReq := fillPixelRequest{
		X: x,
		Y: y,
		RGB: []int{
			constrainToByte(r),
			constrainToByte(g),
			constrainToByte(b),
		},
	}

	var resp Status
	c.httpPut(path, &pixelReq, &resp)

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to set pixel on model %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) GetFonts() (fonts Fonts, err error) {
	if err = c.httpGet("/api/overlays/fonts", &fonts); err != nil {
		return nil, fmt.Errorf("unable to retrieve fonts: %w", err)
	}

	return fonts, err
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

type newArg func(c *Client)

func WithHTTPClient(httpClient *http.Client) newArg {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
