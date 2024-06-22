package main

import (
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/handle"
)

func main() {
	opts := argparse.Parse()
	handle.Handle(opts)
}
