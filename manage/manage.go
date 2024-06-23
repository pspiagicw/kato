package manage

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/player"
)

var titleStyle = lipgloss.NewStyle().Bold(true).Italic(false)

var artistStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
	Light: "#909090",
	Dark:  "#626262",
})

// var artistStyle = lipgloss.NewStyle()
var albumStyle = lipgloss.NewStyle().Italic(false)

func Shuffle(opts *argparse.Opts) {
	player := player.New(opts)

	player.Shuffle()
}

func Toggle(opts *argparse.Opts) {
	player := player.New(opts)

	player.Toggle()
}
func Next(opts *argparse.Opts) {
	player := player.New(opts)

	player.Next()
}
func Prev(opts *argparse.Opts) {
	player := player.New(opts)

	player.Prev()
}

func Play(opts *argparse.Opts) {
	player := player.New(opts)

	player.Play()
}
func Pause(opts *argparse.Opts) {
	player := player.New(opts)

	player.Pause()
}
func Status(opts *argparse.Opts) {
	player := player.New(opts)

	song, err := player.Song()

	if err != nil {
		goreland.LogFatal("Error getting song: %s", err)
	}
	prettyPrint(song)
}
func prettyPrint(song *player.Song) {
	status := "[paused]"
	if song.IsPlaying {
		status = "[playing]"
	}
	title := titleStyle.Render(song.Title)
	artist := artistStyle.Render(song.Artist)

	album := albumStyle.Render(song.Album)
	albumArtist := artistStyle.Render(song.AlbumArtist)
	format := artistStyle.Render(fmt.Sprintf("%s bits ◊ %.1f KHz", song.Bits, float64(song.Frequency)/1000))

	fmt.Printf("%s • %s\n", status, format)
	fmt.Printf("%s • %s\n", title, artist)
	fmt.Printf("%s • %s\n", album, albumArtist)

}
