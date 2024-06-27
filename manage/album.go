package manage

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fhs/gompd/v2/mpd"
	"github.com/gerow/pager"
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/player"
)

func Repeat(opts *argparse.Opts) {
	player := player.New(opts)
	player.Single(true)
	player.Repeat(true)

}

func Playlist(opts *argparse.Opts) {
	player := player.New(opts)
	playlist := player.Playlist()
	current := player.Position()
	next := player.NextPosition()

	printPlaylist(playlist, current, next)
}
func printPlaylist(playlist []mpd.Attrs, current, next int) {
	numbers := make([]string, len(playlist))
	prefix := "  "
	if len(playlist) > 10 {
		pager.Open()
		defer pager.Close()
	}

	for i := range playlist {

		if i == current {
			prefix = "=>"
		} else if i == next {
			prefix = ">>"
		} else {
			prefix = "  "
		}

		numbers[i] = fmt.Sprintf("%s %02d ", prefix, i)
	}

	titles := make([]string, len(playlist))

	for i, song := range playlist {
		titles[i] = fmt.Sprintf("%s ", strings.TrimSpace(song["Title"]))
	}

	artists := make([]string, len(playlist))

	for i, song := range playlist {
		artists[i] = fmt.Sprintf("%s", song["Artist"])
	}

	numberStr := strings.Join(numbers, "\n")
	titleStr := strings.Join(titles, "\n")
	artistStr := strings.Join(artists, "\n")

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, numberStr, titleStr, artistStr))
}

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
