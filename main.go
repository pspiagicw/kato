package main

import (
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/handle"
)

var VERSION = "unversioned"

func main() {
	opts := argparse.Parse(VERSION)
	handle.Handle(opts)
}
