package manage

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fhs/gompd/v2/mpd"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/player"
)

type DashModel struct {
	status player.Song
	player *player.Player

	help help.Model
	keys dashKeyMap
}

type dashKeyMap struct {
	Quit key.Binding
	Next key.Binding
	Prev key.Binding
}

func (d dashKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.Quit, d.Next, d.Prev},
	}
}
func (d dashKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{d.Quit, d.Next, d.Prev}
}

func InitDashModel(opts *argparse.Opts) *DashModel {
	keymap := getDashKeyMap()
	player := player.New(opts)
	status := player.Song()

	return &DashModel{
		keys:   keymap,
		help:   help.New(),
		player: player,
		status: *status,
	}
}
func getDashKeyMap() dashKeyMap {
	return dashKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "Quit"),
		),
		Next: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "Next"),
		),
		Prev: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "Previous"),
		),
	}
}

func (d DashModel) Init() tea.Cmd {
	return nil
}

func (d DashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, d.keys.Quit) {
			return d, tea.Quit
		}
		if key.Matches(msg, d.keys.Next) {
			// d.player.Next()
			return d.nextSong(), nil
		}
		if key.Matches(msg, d.keys.Prev) {
			// d.player.Prev()
			return d.prevSong(), nil
		}
	case tea.WindowSizeMsg:
		d.help.Width = msg.Width
	case statusMsg:
		d.status = msg.status
	}

	return d, nil
}
func (d DashModel) nextSong() tea.Model {
	d.player.Next()
	return d
}
func (d DashModel) prevSong() tea.Model {
	d.player.Prev()
	return d
}

func (d DashModel) View() string {
	song := prettyPrint(&d.status)
	help := d.help.ShortHelpView(d.keys.ShortHelp())
	content := lipgloss.JoinVertical(lipgloss.Left, song, help)
	return content
}
func updateStatus(p *tea.Program, opts *argparse.Opts) mpd.Attrs {
	player := player.New(opts)

	for {
		player.Ping()
		status := player.Song()
		p.Send(statusMsg{*status})
	}
}

type statusMsg struct {
	status player.Song
}

func Dash(opts *argparse.Opts) {
	dash := InitDashModel(opts)
	p := tea.NewProgram(dash)
	go updateStatus(p, opts)
	_, err := p.Run()

	if err != nil {
		// panic(err)
		goreland.LogFatal("Error spawning dash: %v", err)
	}
}
