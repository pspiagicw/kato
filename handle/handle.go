package handle

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/help"
	"github.com/pspiagicw/kato/manage"
)

var handlers = map[string]func(*argparse.Opts){
	"volume":  manage.Volume,
	"play":    manage.Play,
	"pause":   manage.Pause,
	"next":    manage.Next,
	"prev":    manage.Prev,
	"status":  manage.Status,
	"toggle":  manage.Toggle,
	"album":   manage.Album,
	"artist":  manage.Artist,
	"shuffle": manage.Shuffle,
	"stop":    manage.Stop,
	"mark":    notImplemented,
	"save":    notImplemented,
	"version": func(opts *argparse.Opts) {
		help.Version(opts.Version)
	},
	"seek":     manage.Seek,
	"load":     notImplemented,
	"like":     notImplemented,
	"playlist": manage.Playlist,
	"repeat":   manage.Repeat,
	"help": func(opts *argparse.Opts) {
		help.Help()
	},
}

func Handle(opts *argparse.Opts) {
	cmd := opts.Command()

	if cmd == "" {
		help.Help()
		goreland.LogFatal("No command specified")
	}

	handler, ok := handlers[cmd]

	if !ok {
		help.Help()
		goreland.LogFatal("No command '%s'", cmd)
	}

	handler(opts)
}

func notImplemented(*argparse.Opts) {
	goreland.LogFatal("Command not implemented")
}
