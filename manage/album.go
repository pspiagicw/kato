package manage

import (
	"fmt"

	"github.com/fhs/gompd/v2/mpd"
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

func Artist(opts *argparse.Opts) {
	player := player.New(opts)

	artists, err := player.Artists()

	if err != nil {
		goreland.LogFatal("Error querying for artists: %v", err)
	}

	selection := promptSelection("", artists)

	songs, err := player.SongsByArtist(artists[selection])

	if err != nil {
		goreland.LogFatal("Error querying for albums by artist: %v", err)
	}

	// fmt.Println(songs)
	//
	selections := promptMultiple("", getNames(songs))

	selectedSongs := make([]mpd.Attrs, len(selections))

	for i, selection := range selections {
		selectedSongs[i] = songs[selection]
	}

	err = player.PlaySongs(selectedSongs)

	if err != nil {
		goreland.LogFatal("Error playing songs: %v", err)
	}

}
func getNames(songs []mpd.Attrs) []string {
	names := make([]string, len(songs))

	for i, song := range songs {
		names[i] = fmt.Sprintf("%s", song["Title"])
	}

	return names
}
