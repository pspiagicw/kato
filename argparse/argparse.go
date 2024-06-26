package argparse

import (
	"flag"

	"github.com/pspiagicw/yamc/help"
)

type Opts struct {
	Host string
	Port string

	Args    []string
	Version string
}

func (o *Opts) Command() string {
	if len(o.Args) == 0 {
		return ""
	}

	cmd := o.Args[0]
	o.Args = o.Args[1:]

	return cmd

}

func Parse(VERSION string) *Opts {

	flag.Usage = help.Help
	o := new(Opts)

	flag.StringVar(&o.Host, `host`, "127.0.0.1", "Host to connect to")
	flag.StringVar(&o.Port, "port", "6600", "Port to connect to")
	o.Version = VERSION

	flag.Parse()
	o.Args = flag.Args()

	return o
}
