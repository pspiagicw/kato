package player

import (
	"fmt"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/pspiagicw/goreland"
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

func New() *Player {
	p := &Player{}
	p.connect()
	return p

}
func (p *Player) Title() (string, error) {
	attrs, err := p.client.CurrentSong()
	if err != nil {
		return "", err
	}

	return attrs["Title"], nil
}
func (p *Player) Artist() (string, error) {
	attrs, err := p.client.CurrentSong()
	if err != nil {
		return "", err
	}

	return attrs["Artist"], nil
}
func (p *Player) Album() (string, error) {
	attrs, err := p.client.CurrentSong()
	if err != nil {
		return "", err
	}

	return attrs["Album"], nil
}
func (p *Player) AlbumArtist() (string, error) {
	attrs, err := p.client.CurrentSong()
	if err != nil {
		return "", err
	}

	return attrs["AlbumArtist"], nil
}
func (p Player) Song() (*Song, error) {
	title, err := p.Title()
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	artist, err := p.Artist()
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	album, err := p.Album()
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	albumArtist, err := p.AlbumArtist()
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	return &Song{
		Title:       title,
		Artist:      artist,
		Album:       album,
		AlbumArtist: albumArtist,
	}, nil
}

func (p *Player) connect() {

	client, err := mpd.Dial("tcp", "192.168.1.45:6600")

	if err != nil {
		goreland.LogFatal("Failed to connect to MPD server: %v", err)
	}
	p.client = client
}
