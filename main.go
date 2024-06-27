package main

import (
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/handle"
)

var VERSION = "unversioned"

func main() {
	opts := argparse.Parse(VERSION)
	handle.Handle(opts)
}
