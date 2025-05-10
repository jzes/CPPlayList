package startmenu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jzes/CPPlayList/screen"
	"github.com/jzes/CPPlayList/state"
)

const (
	SubTitle       = "--------Menu--------\n\n"
	menuViewFormat = "%s - %s\n"
)

type Handler struct {
	CurrentMenuIndex int
	MenuSumary       []state.State
}

func NewHandler() *Handler {
	return &Handler{
		0,
		[]state.State{
			state.Menu,
			state.LoginDeezer,
			state.LoginSpotify,
		},
	}
}

func (h *Handler) Update(msg tea.KeyMsg) (state.State, tea.Cmd) {
	switch msg.Type {
	case tea.KeyUp:
		h.selectAboveMenu()
	case tea.KeyDown:
		h.selectBelowMenu()
	case tea.KeyEnter:
		return h.currentMenu(), nil
	}

	return h.MenuSumary[0], nil
}

func (h *Handler) selectAboveMenu() {
	if h.CurrentMenuIndex > 0 {
		h.CurrentMenuIndex--
	}
}

func (h *Handler) selectBelowMenu() {
	if h.CurrentMenuIndex < len(h.MenuSumary)-1 {
		h.CurrentMenuIndex++
	}
}

func (h *Handler) currentMenu() state.State {
	return h.MenuSumary[h.CurrentMenuIndex]
}

func (h *Handler) View() string {
	view := screen.Headline
	view += SubTitle

	for i := range h.MenuSumary {
		view += h.printMenu(i)
	}

	return view
}

func (h *Handler) printMenu(menuIndex int) string {
	return fmt.Sprintf(
		menuViewFormat,
		h.getCursor(menuIndex),
		string(h.MenuSumary[menuIndex]),
	)
}

func (h *Handler) getCursor(menuIndex int) string {
	if menuIndex == h.CurrentMenuIndex {
		return screen.SelectedCursor
	}
	return screen.EmptyCursor
}
