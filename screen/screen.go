package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jzes/CPPlayList/state"
)

const (
	Headline       = "------CPPlayList-------\n"
	EmptyCursor    = " "
	SelectedCursor = ">"
)

type Screen struct {
	currentState state.State
	handlers     map[state.State]state.Handler
}

func New(handlers map[state.State]state.Handler) Screen {
	return Screen{
		currentState: state.Menu,
		handlers:     handlers,
	}
}

func (s Screen) Init() tea.Cmd {
	return nil
}

func (s Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgT := msg.(type) {
	case tea.KeyMsg:
		if msgT.Type == tea.KeyEsc {
			return s, tea.Quit
		}

		newState, cmd := s.handlers[s.currentState].Update(msgT)
		s.currentState = newState
		return s, cmd
	}
	return s, nil
}

func (s Screen) View() string {
	return s.handlers[s.currentState].View()
}
