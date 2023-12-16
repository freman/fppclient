package fppclient

import (
	"context"
	"fmt"
)

func (c Client) GetPlaylists(ctx context.Context) (playlists []string, err error) {

	if err = c.httpGet(ctx, "/api/playlists", &playlists); err != nil {
		return nil, fmt.Errorf("unable to retrieve playlists: %w", err)
	}

	return playlists, err
}
