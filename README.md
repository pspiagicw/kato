# kato

`kato` is a MPD client.

- [kato](#kato)
    - [features](#features)
    - [installation](#installation)
    - [documentation](#documentation)
    - [commands](#commands)

# Features

`kato` is a [mpd](https://musicpd.org) client.
It acts like a remote controller to a already running music server.

- Psuedo-TUI like interface.
- Easy & Simple commands.
- Intuitive interface

# Installation

You need to have a configured and running `MPD` server.

You can download a binary from the [releases](https://github.com/psipagicw/kato/releases) section.

Or else if you have the `Go` compiler installed. You can run

```sh  {linenos=false}
go install github.com/pspiagicw/kato@latest
```

If you use [`gox`](https://github.com/pspiagicw/gox), you can also run

```sh {linenos=false}
gox install github.com/pspiagicw/kato@latest
```

To build the project, download the project and run `go build .`

```sh {linenos=false}
git clone https://github.com/pspiagicw/kato
cd kato
go build .
```

# Documentation

`kato` includes a thorough help section.

Just run `kato help`

```sh {linenos=false}
kato help
```

For more help per command, run `kato help [command]`


# Commands

### help

- Simple command to provide help documentation.

### play

- Play the current song, or resume playback.

### pause

- Pause the current song.

### 


