package manage

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/yamc/argparse"
	"github.com/pspiagicw/yamc/player"
)

type VolumeModel struct {
	volume int
	widget progress.Model
	help   help.Model
	keys   volumeKeyMap

	player *player.Player
}

type volumeKeyMap struct {
	Quit  key.Binding
	Raise key.Binding
	Lower key.Binding
}

func (k volumeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Raise, k.Lower},
	}
}
func (k volumeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Raise, k.Lower}
}

func (m VolumeModel) Init() tea.Cmd {
	return nil
}
func NewVolumeModel(opts *argparse.Opts) VolumeModel {
	p := player.New(opts)
	vol := p.Volume()

	keys := getVolumeKeyMap()
	return VolumeModel{
		volume: vol,
		widget: progress.New(),
		keys:   keys,
		help:   help.New(),
		player: p,
	}
}
func getVolumeKeyMap() volumeKeyMap {
	return volumeKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "Quit"),
		),
		Raise: key.NewBinding(
			key.WithKeys("+", "k", "="),
			key.WithHelp("+", "Raise volume"),
		),
		Lower: key.NewBinding(
			key.WithKeys("-", "j"),
			key.WithHelp("-", "Lower volume"),
		),
	}
}
func (m VolumeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.widget.Width = msg.Width * 2 / 3
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
		if key.Matches(msg, m.keys.Raise) {
			return m.raiseVolume(), nil
		}
		if key.Matches(msg, m.keys.Lower) {
			return m.lowerVolume(), nil
		}
	}
	return m, nil
}
func (m VolumeModel) raiseVolume() VolumeModel {
	m.volume += 2

	m.volume = min(100, m.volume)

	m.player.SetVolume(m.volume)

	return m

}
func (m VolumeModel) lowerVolume() VolumeModel {
	m.volume -= 2

	m.volume = max(0, m.volume)

	m.player.SetVolume(m.volume)

	return m

}
func (m VolumeModel) View() string {
	percent := float64(m.volume) / 100

	bar := m.widget.ViewAs(percent)
	help := m.help.ShortHelpView(m.keys.ShortHelp())

	content := lipgloss.JoinVertical(lipgloss.Left, bar, help)

	padding := lipgloss.NewStyle().PaddingTop(1).PaddingLeft(2).PaddingBottom(1)
	return padding.Render(content)
}

func Volume(opts *argparse.Opts) {
	p := tea.NewProgram(NewVolumeModel(opts))
	_, err := p.Run()
	if err != nil {
		// panic(err)
		goreland.LogError("Error spawning volume: %v", err)
	}
}
