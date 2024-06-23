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

func selectArtist(player *player.Player) string {
	artists := player.Artists()

	selection := promptSelection("", artists)

	return artists[selection]
}
func selectSongs(player *player.Player, artist string) []mpd.Attrs {
	songs := player.SongsByArtist(artist)

	selections := promptMultiple("", getNames(songs))

	selectedSongs := make([]mpd.Attrs, len(selections))

	for i, selection := range selections {
		selectedSongs[i] = songs[selection]
	}

	return selectedSongs
}
func Artist(opts *argparse.Opts) {
	player := player.New(opts)

	artist := selectArtist(player)

	songs := selectSongs(player, artist)

	player.PlaySongs(songs)
}
func getNames(songs []mpd.Attrs) []string {
	names := make([]string, len(songs))

	for i, song := range songs {
		names[i] = fmt.Sprintf("%s", song["Title"])
	}

	return names
}
