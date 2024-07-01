package player

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/kato/argparse"
)

type Player struct {
	client *mpd.Client
}

func New(opts *argparse.Opts) *Player {
	p := &Player{}
	p.connect(opts.Host, opts.Port)
	return p

}
func (p *Player) Ping() {
	p.client.Ping()
}
func (p *Player) SeekTo(elapsed float64) {
	seekTime := time.Second * time.Duration(elapsed)

	err := p.client.SeekCur(seekTime, false)

	if err != nil {
		goreland.LogFatal("Failed to seek to position: %v", err)
	}
}

func (p *Player) Repeat(state bool) {
	err := p.client.Repeat(state)

	if err != nil {
		goreland.LogFatal("Failed to set repeat state: %v", err)
	}
}
func (p *Player) Single(state bool) {
	err := p.client.Single(state)

	if err != nil {
		goreland.LogFatal("Failed to set single state: %v", err)
	}
}
func (p *Player) Stop() {
	err := p.client.Stop()

	if err != nil {
		goreland.LogFatal("Failed to stop playback: %v", err)
	}
}
func (p *Player) Clear() {
	err := p.client.Clear()

	if err != nil {
		goreland.LogFatal("Failed to clear playlist: %v", err)
	}

}
func (p *Player) Random(state bool) {
	err := p.client.Random(state)

	if err != nil {
		goreland.LogFatal("Failed to set random state: %v", err)
	}

}
func (p *Player) Shuffle() {

	p.Clear()
	p.Random(true)

	err := p.client.Add("/")

	if err != nil {
		goreland.LogFatal("Failed to add all songs to playlist: %v", err)
	}

	p.Play()
}
func (p *Player) Add(track mpd.Attrs) {
	err := p.client.Add(track["file"])

	if err != nil {
		goreland.LogFatal("Failed to add song to playlist: %v", err)
	}

}
func (p *Player) PlaySongs(songs []mpd.Attrs) {
	p.Clear()

	for _, track := range songs {
		p.Add(track)
	}

	p.Play()
}
func (p *Player) PlayAlbum(album string) {

	p.Random(false)

	tracks := p.SongsByAlbum(album)

	p.PlaySongs(tracks)
}

func (p *Player) IsPlaying() bool {
	status := p.Status()
	return status["state"] != "pause"
}

func (p *Player) Toggle() {
	if p.IsPlaying() {
		p.Pause(true)
	} else {
		p.Pause(false)

	}
}
func (p *Player) SetVolume(vol int) {
	err := p.client.SetVolume(vol)

	if err != nil {
		goreland.LogFatal("Failed to set volume: %v", err)
	}
}
func (p *Player) Next() {
	err := p.client.Next()

	if err != nil {
		goreland.LogFatal("Failed to play next: %v", err)
	}
}
func (p *Player) Prev() {
	err := p.client.Previous()

	if err != nil {
		goreland.LogFatal("Failed to play previous: %v", err)
	}
}
func (p *Player) Play() {
	err := p.client.Play(-1)

	if err != nil {
		goreland.LogFatal("Failed to play: %v", err)
	}
}
func (p *Player) Pause(status bool) {
	err := p.client.Pause(status)

	if err != nil {
		goreland.LogFatal("Failed to pause: %v", err)
	}
}
func (p *Player) Volume() int {
	status := p.Status()

	volStr := status["volume"]

	return toInt(volStr)

}
func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		goreland.LogFatal("Failed to convert string to int: %v", err)
	}
	return i
}

func (p *Player) Status() mpd.Attrs {
	status, err := p.client.Status()
	if err != nil {
		goreland.LogFatal("Failed to get status: %v", err)
	}
	return status
}

func (p *Player) connect(host, port string) {

	addr := fmt.Sprintf("%s:%s", host, port)

	client, err := mpd.Dial("tcp", addr)

	if err != nil {
		goreland.LogFatal("Failed to connect to MPD server: %v", err)
	}
	p.client = client
}
