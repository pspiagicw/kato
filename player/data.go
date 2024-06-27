package player

import (
	"strconv"
	"strings"

	"github.com/fhs/gompd/v2/mpd"
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

func (p *Player) Seek() (float64, float64) {
	status := p.Status()
	elapsed, err := strconv.ParseFloat(status["elapsed"], 32)

	if err != nil {
		goreland.LogFatal("Failed to get elapsed time: %v", err)
	}

	total, err := strconv.ParseFloat(status["duration"], 32)

	if err != nil {
		goreland.LogFatal("Failed to get total time: %v", err)
	}

	return elapsed, total
}

func (p *Player) Playlist() []mpd.Attrs {
	tracks, err := p.client.PlaylistInfo(-1, -1)
	if err != nil {
		goreland.LogFatal("Failed to get playlist: %v", err)
	}

	return tracks
}
func (p *Player) Position() int {
	pos, err := p.client.Status()
	if err != nil {
		goreland.LogFatal("Failed to get current song position: %v", err)
	}
	result, err := strconv.Atoi(pos["song"])

	if err != nil {
		goreland.LogFatal("Failed to convert song position to integer: %v", err)
	}

	return result
}
func (p *Player) NextPosition() int {
	pos, err := p.client.Status()
	if err != nil {
		goreland.LogFatal("Failed to get current song position: %v", err)
	}
	result, err := strconv.Atoi(pos["nextsong"])

	if err != nil {
		goreland.LogFatal("Failed to convert song position to integer: %v", err)
	}

	return result
}

func (p *Player) SongsByArtist(artist string) []mpd.Attrs {
	tracks, err := p.client.Search("artist", artist)

	if err != nil {
		goreland.LogFatal("Failed to find tracks by artist: %v", err)
	}

	return tracks
}

func (p *Player) SongsByAlbum(album string) []mpd.Attrs {
	tracks, err := p.client.Find("album", album)

	if err != nil {
		goreland.LogFatal("Failed to find tracks by album: %v", err)
	}

	return tracks
}

func (p Player) Artists() []string {
	tracks, err := p.client.List("artist")

	if err != nil {
		goreland.LogFatal("Failed to list artists: %v", err)
	}

	return tracks
}

func (p *Player) Albums() []string {
	albums, err := p.client.List("album")
	if err != nil {
		goreland.LogFatal("Failed to list albums: %v", err)
	}
	return albums
}

func (p Player) Song() *Song {
	attrs, err := p.client.CurrentSong()

	if err != nil {
		goreland.LogFatal("Failed to get current song: %v", err)
	}

	isPlaying := p.IsPlaying()

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
	}
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
