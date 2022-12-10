package fppclient

import (
	"fmt"
	"strings"
)

func (c Client) GetOverlaysModels() (models Models, err error) {
	if err = c.httpGet("/api/overlays/models", &models); err != nil {
		return nil, fmt.Errorf("unable to retrieve models: %w", err)
	}

	return models, err
}

func (c Client) GetOverlaysModel(name string) (*Model, error) {
	var model Model

	path := fmt.Sprintf("/api/overlays/model/%s", name)
	if err := c.httpGet(path, &model); err != nil {
		return nil, fmt.Errorf("unable to retrieve model %q: %w", name, err)
	}

	return &model, nil
}

func (c Client) ClearOverlaysModel(name string) error {
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

func (c Client) GetOverlaysModelData(name string, rle bool) (*ModelData, error) {
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

func (c Client) SetOverlaysModelState(name string, state interface{}) error {
	path := fmt.Sprintf("/api/overlays/model/%s/state", name)

	var resp Status
	if err := c.httpPut(path, struct {
		State interface{}
	}{State: state}, &resp); err != nil {
		return fmt.Errorf("unable to set model state: %q: %w", name, err)
	}

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to set model state %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) FillOverlaysModel(name string, r, g, b int) error {
	path := fmt.Sprintf("/api/overlays/model/%s/fill", name)

	fillReq := fillPixelRequest{
		RGB: []int{
			constrainToByte(r),
			constrainToByte(g),
			constrainToByte(b),
		},
	}

	var resp Status
	if err := c.httpPut(path, &fillReq, &resp); err != nil {
		return fmt.Errorf("unable to fill model %q: %err", name, err)
	}

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to fill model %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) SetOverlaysModelPixel(name string, x, y, r, g, b int) error {
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
	if err := c.httpPut(path, &pixelReq, &resp); err != nil {
		return fmt.Errorf("unable to set pixel on model %q: %w", name, err)
	}

	if !strings.EqualFold(resp.Status, "OK") {
		return fmt.Errorf("unable to set pixel on model %q: %s", name, resp.Message)
	}

	return nil
}

func (c Client) GetOverlaysFonts() (fonts Fonts, err error) {
	if err = c.httpGet("/api/overlays/fonts", &fonts); err != nil {
		return nil, fmt.Errorf("unable to retrieve fonts: %w", err)
	}

	return fonts, err
}
