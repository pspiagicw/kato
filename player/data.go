package player

import "fmt"

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
