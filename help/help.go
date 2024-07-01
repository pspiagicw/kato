package help

import "github.com/pspiagicw/pelp"

func Version(version string) {
	pelp.Version("kato", version)
}
func Help() {
	pelp.Print("Yet another MPD client")
	pelp.HeaderWithDescription("Usage", []string{"kato [flags] [commands]"})
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
	pelp.Aligned(
		"commands",
		[]string{
			"volume",
			"play",
			"pause",
			"next",
			"prev",
			"status",
			"toggle",
			"album",
			"artist",
			"shuffle",
			"stop",
			"version",
			"seek",
			"playlist",
			"repeat",
			"dash",
			"help",
		},
		[]string{
			"Manage the volume",
			"Start playing",
			"Pause playback",
			"Play the next song",
			"Play the previous song",
			"Get the current status",
			"Toggle play/pause",
			"Play an album",
			"Play specific songs of an artist",
			"Run all songs in a random order",
			"Stop playback",
			"Show the version",
			"Seek to a specific time",
			"Show the current playlist",
			"Repeat current song",
			"Show the dashboard",
			"Show this help message",
		},
	)
	pelp.HeaderWithDescription(
		"More Help",
		[]string{
			"Run 'kato help [command]' for more information on a command",
			"Visit https://github.com/pspiagicw/kato for more information",
		},
	)
}
