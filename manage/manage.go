package manage

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/player"
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
func Stop(opts *argparse.Opts) {
	player := player.New(opts)

	player.Stop()
}
func Pause(opts *argparse.Opts) {
	player := player.New(opts)

	player.Pause(true)
}
func Status(opts *argparse.Opts) {
	player := player.New(opts)

	song := player.Song()

	output := prettyPrint(song)
	fmt.Print(output)
}
func prettyPrint(song *player.Song) string {

	var buffer strings.Builder
	status := "[paused]"
	if song.IsPlaying {
		status = "[playing]"
	}
	title := titleStyle.Render(song.Title)
	artist := artistStyle.Render(song.Artist)

	album := albumStyle.Render(song.Album)
	albumArtist := artistStyle.Render(song.AlbumArtist)
	format := artistStyle.Render(fmt.Sprintf("%s bits ◊ %.1f KHz", song.Bits, float64(song.Frequency)/1000))

	buffer.WriteString(fmt.Sprintf("%s • %s\n", status, format))
	buffer.WriteString(fmt.Sprintf("%s • %s\n", title, artist))
	buffer.WriteString(fmt.Sprintf("%s • %s\n", album, albumArtist))

	return buffer.String()

}
