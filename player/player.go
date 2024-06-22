package player

import (
	"fmt"
	"strconv"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/yamc/argparse"
)

type Player struct {
	client *mpd.Client
}

type Song struct {
	Title       string
	Artist      string
	Album       string
	AlbumArtist string
}

func New(opts *argparse.Opts) *Player {
	p := &Player{}
	p.connect(opts.Host, opts.Port)
	return p

}

func (p *Player) Toggle() error {
	status, err := p.client.Status()
	if err != nil {
		return err
	}
	if status["state"] == "pause" {
		return p.client.Pause(false)
	} else {
		return p.client.Pause(true)
	}
}
func (p *Player) SetVolume(vol int) error {
	return p.client.SetVolume(vol)
}
func (p *Player) Next() error {
	return p.client.Next()
}
func (p *Player) Prev() error {
	return p.client.Previous()
}
func (p *Player) Play() error {
	return p.client.Play(-1)
}
func (p *Player) Pause() error {
	return p.client.Pause(true)
}
func (p *Player) Volume() (int, error) {
	status, err := p.client.Status()
	if err != nil {
		return 0, err
	}
	volStr := status["volume"]

	vol, err := strconv.ParseInt(volStr, 10, 32)

	if err != nil {
		return 0, err
	}
	return int(vol), nil
}

func (p *Player) connect(host, port string) {

	addr := fmt.Sprintf("%s:%s", host, port)

	client, err := mpd.Dial("tcp", addr)

	if err != nil {
		goreland.LogFatal("Failed to connect to MPD server: %v", err)
	}
	p.client = client
}
