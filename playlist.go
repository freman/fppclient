package fppclient

import (
	"context"
	"fmt"
)

func (c Client) GetPlaylist(ctx context.Context, name string) (playlist Playlist, err error) {
	path := fmt.Sprintf("/api/playlist/%s", name)

	if err = c.httpGet(ctx, path, &playlist); err != nil {
		return playlist, fmt.Errorf("unable to retrieve playlists: %w", err)
	}

	return playlist, err
}

type Playlist struct {
	Name         string            `json:"name"`
	Version      int               `json:"version"`
	Repeat       int               `json:"repeat"`
	LoopCount    int               `json:"loopCount"`
	Empty        bool              `json:"empty"`
	Desc         string            `json:"desc"`
	Random       int               `json:"random"`
	LeadIn       []PlaylistEntries `json:"leadIn"`
	MainPlaylist []PlaylistEntries `json:"mainPlaylist"`
	LeadOut      []PlaylistEntries `json:"leadOut"`
	PlaylistInfo PlaylistInfo      `json:"playlistInfo"`
}

type PlaylistEntries struct {
	Type         string  `json:"type"`
	Enabled      int     `json:"enabled"`
	PlayOnce     int     `json:"playOnce"`
	SequenceName string  `json:"sequenceName"`
	MediaName    string  `json:"mediaName"`
	VideoOut     string  `json:"videoOut"`
	Duration     float64 `json:"duration"`
}

type PlaylistInfo struct {
	TotalDuration float64 `json:"total_duration"`
	TotalItems    int     `json:"total_items"`
}
