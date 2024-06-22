package handle

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/help"
	"github.com/pspiagicw/yamc/manage"
)

var handlers = map[string]func(*argparse.Opts){
	"volume": manage.Volume,
	"play":   manage.Play,
	"pause":  manage.Pause,
	"next":   manage.Next,
	"prev":   manage.Prev,
	"status": manage.Status,
	"toggle": manage.Toggle,
	"album":  manage.Album,
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
