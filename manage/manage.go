package manage

import (
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/player"
)

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
