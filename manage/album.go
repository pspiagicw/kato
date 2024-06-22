package manage

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/player"
)

func Album(opts *argparse.Opts) {
	player := player.New(opts)

	albums, err := player.Albums()

	if err != nil {
		goreland.LogFatal("Error querying for albums: %v", err)
	}

	selection := promptSelection("", albums)

	player.PlayAlbum(albums[selection])
}
