package fppclient

import (
	"context"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func (c Client) PostCommand(ctx context.Context, cmd Command) {
	res := map[string]interface{}{}
	if err := c.httpPost(ctx, "/api/command", cmd, res); err != nil {
		panic(err)
	}
	spew.Dump(res)
}

func CommandInsertPlaylistAfterCurrent(playlistName string, startIndex, endIndex int, ifNotRunning bool) Command {
	return Command{
		Command: "Insert Playlist After Current",
		Args: []string{
			playlistName,
			strconv.Itoa(startIndex),
			strconv.Itoa(endIndex),
			strconv.FormatBool(ifNotRunning),
		},
	}
}
