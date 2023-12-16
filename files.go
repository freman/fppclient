package fppclient

import (
	"context"
	"fmt"
)

func (c Client) GetFiles(ctx context.Context, dir string) (files []File, err error) {
	var f Files
	path := fmt.Sprintf("/api/files/%s", dir)
	if err = c.httpGet(ctx, path, &f); err != nil {
		return nil, fmt.Errorf("unable to retrieve files: %w", err)
	}

	return f.Files, err
}

type Files struct {
	Status string `json:"status"`
	Files  []File `json:"files"`
}
type File struct {
	Name      string `json:"name"`
	Mtime     string `json:"mtime"`
	SizeBytes int    `json:"sizeBytes"`
	SizeHuman string `json:"sizeHuman"`
}
