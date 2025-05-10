package deezer

import (
	"github.com/jzes/CPPlayList/screen"
	"github.com/jzes/CPPlayList/state"
)

const (
	subtitle        = "🟣-Credenciais Deezer-\n\n🔐 Digite as credenciais do deezer\n"
	fieldViewFormat = "%s %s\n╰─❯%s\n"
)

type LoginHandler struct {
	*screen.Form
}

func NewLoginHandler() *LoginHandler {
	f := &screen.Form{
		SelfState:       state.LoginDeezer,
		Subtitle:        subtitle,
		FieldViewFormat: fieldViewFormat,
		FieldCursor:     0,
		Fields:          []string{"user"},
		Values:          make([]string, 1),
	}
	return &LoginHandler{
		Form: f,
	}
}
