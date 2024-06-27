package manage

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/kato/argparse"
	"github.com/pspiagicw/kato/player"
)

type seekMsg struct {
	elapsed float64
	total   float64
}

type SeekModel struct {
	elapsed float64
	total   float64
	widget  progress.Model
	help    help.Model
	keys    seekKeyMap
	player  *player.Player
}

type seekKeyMap struct {
	Quit     key.Binding
	Forward  key.Binding
	Backward key.Binding
}

func (s seekKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{s.Quit, s.Forward, s.Backward},
	}
}
func (s seekKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{s.Quit, s.Forward, s.Backward}
}

func (m SeekModel) Init() tea.Cmd {
	return nil
}

func NewSeekModel(opts *argparse.Opts) SeekModel {

	p := player.New(opts)
	keys := getSeekKeyMap()
	elapsed, total := p.Seek()

	return SeekModel{
		elapsed: elapsed,
		total:   total,
		widget:  progress.New(),
		keys:    keys,
		help:    help.New(),
		player:  p,
	}
}
func getSeekKeyMap() seekKeyMap {
	return seekKeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "Quit"),
		),
		Forward: key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("l", "Forward"),
		),
		Backward: key.NewBinding(
			key.WithKeys("h", "left"),
			key.WithHelp("h", "Backward"),
		),
	}
}

func (s SeekModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.widget.Width = msg.Width * 2 / 3
	case tea.KeyMsg:
		if key.Matches(msg, s.keys.Quit) {
			return s, tea.Quit
		}
		if key.Matches(msg, s.keys.Forward) {
			return s.seekForward(), nil
		}
		if key.Matches(msg, s.keys.Backward) {
			return s.seekBackward(), nil
		}
	case seekMsg:
		s.elapsed = msg.elapsed
		s.total = msg.total
	}
	return s, nil
}
func (s SeekModel) seekForward() tea.Model {
	s.elapsed += 2

	s.elapsed = min(s.total, s.elapsed)

	s.player.SeekTo(s.elapsed)
	return s
}
func (s SeekModel) seekBackward() tea.Model {
	s.elapsed -= 2

	s.elapsed = max(0, s.elapsed)

	s.player.SeekTo(s.elapsed)
	return s
}
func (s SeekModel) View() string {
	percent := s.elapsed / s.total

	bar := s.widget.ViewAs(percent)
	help := s.help.ShortHelpView(s.keys.ShortHelp())

	content := lipgloss.JoinVertical(lipgloss.Left, bar, help)

	padding := lipgloss.NewStyle().PaddingTop(1).PaddingLeft(2).PaddingBottom(1)
	return padding.Render(content)
}
func updateProgress(p *tea.Program, opts *argparse.Opts) {
	player := player.New(opts)
	for {
		elapsed, total := player.Seek()
		p.Send(seekMsg{elapsed, total})
	}
}
func Seek(opts *argparse.Opts) {
	p := tea.NewProgram(NewSeekModel(opts))
	go updateProgress(p, opts)
	_, err := p.Run()

	if err != nil {
		goreland.LogFatal("Failed to run program: %v", err)
	}
}
