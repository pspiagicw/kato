package manage

import (
	"fmt"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/player"
)

func Album(opts *argparse.Opts) {
	player := player.New(opts)

	albums := player.Albums()

	selection := promptSelection("", albums)

	player.PlayAlbum(albums[selection])
}

func Artist(opts *argparse.Opts) {
	player := player.New(opts)

	artists := player.Artists()

	selection := promptSelection("", artists)

	songs := player.SongsByArtist(artists[selection])

	selections := promptMultiple("", getNames(songs))

	selectedSongs := make([]mpd.Attrs, len(selections))

	for i, selection := range selections {
		selectedSongs[i] = songs[selection]
	}

	player.PlaySongs(selectedSongs)
}
func getNames(songs []mpd.Attrs) []string {
	names := make([]string, len(songs))

	for i, song := range songs {
		names[i] = fmt.Sprintf("%s", song["Title"])
	}

	return names
}
