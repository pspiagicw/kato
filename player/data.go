package player

import (
	"strconv"
	"strings"

	"github.com/pspiagicw/goreland"
)

type Song struct {
	Title       string
	Artist      string
	Album       string
	AlbumArtist string
	IsPlaying   bool
	Format      string
	Bits        string
	Frequency   int
}

func (p Player) Song() (*Song, error) {
	attrs, err := p.client.CurrentSong()

	if err != nil {
		return nil, err
	}

	isPlaying, err := p.IsPlaying()

	if err != nil {
		return nil, err
	}

	title := attrs["Title"]
	artist := attrs["Artist"]
	album := attrs["Album"]
	albumArtist := attrs["AlbumArtist"]
	format := attrs["Format"]
	freq, bits := parseFormat(format)

	return &Song{
		Title:       title,
		Artist:      artist,
		Album:       album,
		AlbumArtist: albumArtist,
		Format:      format,
		IsPlaying:   isPlaying,
		Frequency:   freq,
		Bits:        bits,
	}, nil
}
func parseFormat(format string) (int, string) {
	parts := strings.Split(format, ":")
	if len(parts) != 3 {
		return 0, "Unknown"
	}
	freq, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		goreland.LogFatal("Error parsing frequency: %v", err)
	}
	return int(freq), parts[1]
}
