package help

import "github.com/pspiagicw/pelp"

func Help() {
	pelp.Print("Yet another MPD client")
	pelp.HeaderWithDescription("Usage", []string{"yamc [flags] [commands]"})
	pelp.Flags(
		"flags",
		[]string{
			"help",
			"host",
			"port",
		},
		[]string{
			"Show this help message",
			"MPD host to connect to",
			"MPD port to connect to",
		},
	)
	pelp.HeaderWithDescription(
		"More Help",
		[]string{
			"Run 'yamc help [command]' for more information on a command",
			"Visit https://github.com/pspiagicw/yamc for more information",
		},
	)
}
