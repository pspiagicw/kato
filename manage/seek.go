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
	Quit        key.Binding
	Forward     key.Binding
	Backward    key.Binding
	FirstStop   key.Binding
	SecondStop  key.Binding
	ThirdStop   key.Binding
	FourthStop  key.Binding
	FifthStop   key.Binding
	SixthStop   key.Binding
	SeventhStop key.Binding
	EighthStop  key.Binding
	NinthStop   key.Binding
	LastStop    key.Binding
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
		FirstStop: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "Seek to 10%"),
		),
		SecondStop: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "Seek to 20%"),
		),
		ThirdStop: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "Seek to 30%"),
		),
		FourthStop: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("4", "Seek to 40%"),
		),
		FifthStop: key.NewBinding(
			key.WithKeys("5"),
			key.WithHelp("5", "Seek to 50%"),
		),
		SixthStop: key.NewBinding(
			key.WithKeys("6"),
			key.WithHelp("6", "Seek to 60%"),
		),
		SeventhStop: key.NewBinding(
			key.WithKeys("7"),
			key.WithHelp("7", "Seek to 70%"),
		),
		EighthStop: key.NewBinding(
			key.WithKeys("8"),
			key.WithHelp("8", "Seek to 80%"),
		),
		NinthStop: key.NewBinding(
			key.WithKeys("9"),
			key.WithHelp("9", "Seek to 90%"),
		),
		LastStop: key.NewBinding(
			key.WithKeys("0"),
			key.WithHelp("0", "Seek to 100%"),
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
		if key.Matches(msg, s.keys.FirstStop) {
			return s.seekTo(10), nil
		}
		if key.Matches(msg, s.keys.SecondStop) {
			return s.seekTo(20), nil
		}
		if key.Matches(msg, s.keys.ThirdStop) {
			return s.seekTo(30), nil
		}
		if key.Matches(msg, s.keys.FourthStop) {
			return s.seekTo(40), nil
		}
		if key.Matches(msg, s.keys.FifthStop) {
			return s.seekTo(50), nil
		}
		if key.Matches(msg, s.keys.SixthStop) {
			return s.seekTo(60), nil
		}
		if key.Matches(msg, s.keys.SeventhStop) {
			return s.seekTo(70), nil
		}
		if key.Matches(msg, s.keys.EighthStop) {
			return s.seekTo(80), nil
		}
		if key.Matches(msg, s.keys.NinthStop) {
			return s.seekTo(90), nil
		}
		if key.Matches(msg, s.keys.LastStop) {
			return s.seekTo(100), nil
		}
	case seekMsg:
		s.elapsed = msg.elapsed
		s.total = msg.total
	}
	return s, nil
}
func (s SeekModel) seekTo(percent float64) tea.Model {
	s.elapsed = percent * s.total / 100
	s.player.SeekTo(s.elapsed)
	return s
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
