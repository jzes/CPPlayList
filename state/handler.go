package state

import tea "github.com/charmbracelet/bubbletea"

type Handler interface {
	Update(tea.KeyMsg) (State, tea.Cmd)
	View() string
}

type State string

const (
	Menu         = State("menu")
	LoginDeezer  = State("login-deezer")
	LoginSpotify = State("login-spotify")
)
