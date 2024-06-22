package argparse

import "flag"

type Opts struct {
	Host string
	Port string

	Args []string
}

func (o *Opts) Command() string {
	if len(o.Args) == 0 {
		return ""
	}

	cmd := o.Args[0]
	o.Args = o.Args[1:]

	return cmd

}

func Parse() *Opts {
	o := new(Opts)

	flag.StringVar(&o.Host, `host`, "localhost", "Host to connect to")
	flag.StringVar(&o.Port, "port", "6666", "Port to connect to")

	flag.Parse()
	o.Args = flag.Args()

	return o
}
