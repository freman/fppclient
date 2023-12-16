package fppclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BodyError struct {
	body []byte
	err  error
}

func (c Client) formatURL(path string) string {
	return c.baseURL.ResolveReference(
		&url.URL{
			Path: path,
		}).String()
}

func (c Client) httpGet(ctx context.Context, path string, v interface{}) error {
	u := c.formatURL(path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	return c.httpDo(req, v)
}

func (c Client) httpPost(ctx context.Context, path string, in, out interface{}) error {
	return c.httpDoWithJSON(ctx, http.MethodPost, path, in, out)

}

func (c Client) httpPut(ctx context.Context, path string, in, out interface{}) error {
	return c.httpDoWithJSON(ctx, http.MethodPut, path, in, out)
}

func (c Client) httpDoWithJSON(ctx context.Context, method, path string, in, out interface{}) error {
	u := c.formatURL(path)

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(in); err != nil {
		return fmt.Errorf("unable to marshal object: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, &buf)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.httpDo(req, out)
}

func (c Client) httpDo(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Help the server out by reading and discarding the body
		io.Copy(io.Discard, resp.Body) //nolint:errcheck // don't actually care, we're just trying to be nice.
		return fmt.Errorf("unexpected HTTP status %q (%d)", resp.Status, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("unable to parse response: %w", err)
	}

	return nil
}
