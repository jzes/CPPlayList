package screen

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jzes/CPPlayList/state"
)

type Form struct {
	FieldCursor     int
	Fields          []string
	Values          []string
	Subtitle        string
	FieldViewFormat string
	SelfState       state.State
}

func (f *Form) Update(kMsg tea.KeyMsg) (state.State, tea.Cmd) {
	switch kMsg.Type {
	case tea.KeyEnter:
		if f.isDone() {
			return state.Menu, nil
		}
		f.FieldCursor++
	case tea.KeyBackspace:
		f.eraseFromCurrentField()
	default:
		f.writeToCurrentField(kMsg)
	}
	return f.SelfState, nil
}

func (f *Form) isDone() bool {
	return f.FieldCursor >= len(f.Fields)-1
}

func (f *Form) eraseFromCurrentField() {
	currentField := f.Values[f.FieldCursor]
	if len(currentField) > 0 {
		currentField = currentField[:len(currentField)-1]
	}
	f.Values[f.FieldCursor] = currentField
}

func (f *Form) writeToCurrentField(kMsg tea.KeyMsg) {
	f.Values[f.FieldCursor] += kMsg.String()
}

func (f *Form) View() string {
	view := Headline
	view += f.Subtitle

	for i := range f.Fields {
		view += f.sPrintField(i)
	}
	return view
}

func (f *Form) sPrintField(fieldIndex int) string {
	return fmt.Sprintf(
		f.FieldViewFormat,
		f.getCursor(fieldIndex),
		f.Fields[fieldIndex],
		f.Values[fieldIndex],
	)
}

func (f *Form) getCursor(fieldIndex int) string {
	if fieldIndex == f.FieldCursor {
		return SelectedCursor
	}
	return EmptyCursor
}
