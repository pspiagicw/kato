# kato

`kato` is a MPD client.

- [kato](#kato)
    - [features](#features)
    - [installation](#installation)
    - [documentation](#documentation)
    - [commands](#commands)
    - [purpose](#purpose)
    - [contribution](#contribution)

# Features

`kato` is a [mpd](https://musicpd.org) client.
It acts like a remote controller to a already running music server.

- Psuedo-TUI interface.
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

### Simple commands

| Command | Description           |
|---------|-----------------------|
| play    | Play the current song |
| stop    | Stop playback         |
| pause   | Pause playback        |
| next    | Play next song        |
| prev    | Play previous song    |
| toggle  | Toggle playback       |
| version | Print version info    |

### status

- Print info about current song
- Includes title, artist, quant size and frequency

![status](./gifs/status.gif)

### album

- Run a album
- Interactively select the album
- The current queue is cleared.
- `random` is turned off.

![album](./gifs/album.gif)

### artist

- Play songs by a specific artist
- Multi-select songs.
- The current queue is cleared
- `random` is turned off.

![volume](./gifs/artist.gif)

### shuffle

- Shuffle all songs in the library.
- `random` is turned on.

### volume

- Control volume interactively.

![volume](./gifs/volume.gif)

### seek

- Control song position interactively.

![seek](./gifs/seek.gif)

### playlist

- Display the playlist in a pager.

![playlist](./gifs/playlist.gif)

### dash

- Display a simple dash UI with essential info.
- Simple keybinds to next/prev songs.

![dash](./gifs/dash.gif)

### repeat

- Repeat current song indefinately.
- `repeat` is turned on.
- `single` is turned on.

# Purpose

`kato` doesn't aim to replace tools like `ncmpcpp` and `mpc`.

It only provides a helpful interface for common interactions like
- Playing a album
- Play all songs but shuffled.

It provides quality of life features that would require some scripting or a learning curve in other clients.
If you want complex, minute edits you WILL need to use tools like `mpc`.

Or if you need a TUI experience complete with a visualizer and other goodies, you will need to use `ncmpcpp`.

# Contribution

This tool is made to provide qol features, thus your opinion and contribution is essential to this project.

- You can file a issue to request a feature or report a bug
- If you a little bit of Go, you can fix/implement said feature/bug and file a PR.

You can also show your love for `kato` by giving this project a star on [`GitHub`](https://github.com/pspiagicw/kato).


