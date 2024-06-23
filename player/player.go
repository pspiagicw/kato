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

func New(opts *argparse.Opts) *Player {
	p := &Player{}
	p.connect(opts.Host, opts.Port)
	return p

}
func (p *Player) PlaySongs(songs []mpd.Attrs) error {
	err := p.client.Clear()

	if err != nil {
		return err
	}

	for _, track := range songs {
		err := p.client.Add(track["file"])

		if err != nil {
			return err
		}
	}

	return p.Play()
}
func (p *Player) PlayAlbum(album string) error {

	err := p.client.Random(false)

	if err != nil {
		return err
	}

	tracks, err := p.client.Find("album", album)

	if err != nil {
		return err
	}

	return p.PlaySongs(tracks)
}

func (p *Player) IsPlaying() (bool, error) {
	status, err := p.client.Status()
	if err != nil {
		return false, err
	}
	return status["state"] != "pause", nil
}

func (p *Player) Toggle() error {
	isPlaying, err := p.IsPlaying()
	if err != nil {
		return err
	}
	if isPlaying {
		return p.client.Pause(true)
	} else {
		return p.client.Pause(false)
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
func (p *Player) Status() (map[string]string, error) {
	status, err := p.client.Status()
	if err != nil {
		return map[string]string{}, err
	}
	return status, nil
}

func (p *Player) connect(host, port string) {

	addr := fmt.Sprintf("%s:%s", host, port)

	client, err := mpd.Dial("tcp", addr)

	if err != nil {
		goreland.LogFatal("Failed to connect to MPD server: %v", err)
	}
	p.client = client
}
