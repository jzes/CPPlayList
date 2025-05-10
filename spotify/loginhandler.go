package spotify

import (
	"github.com/jzes/CPPlayList/screen"
	"github.com/jzes/CPPlayList/state"
)

const (
	subtitle        = "🟢-Credenciais Spotify-\n\n🔐 Digite as credenciais do spotify\n"
	fieldViewFormat = "%s %s\n╰─❯%s\n"
)

type LoginHandler struct {
	*screen.Form
}

func NewLoginHandler() *LoginHandler {
	f := &screen.Form{
		SelfState:       state.LoginSpotify,
		Subtitle:        subtitle,
		FieldViewFormat: fieldViewFormat,
		FieldCursor:     0,
		Fields:          []string{"user id", "user secrete"},
		Values:          make([]string, 2),
	}
	return &LoginHandler{
		Form: f,
	}
}
