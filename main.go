package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jzes/CPPlayList/deezer"
	"github.com/jzes/CPPlayList/screen"
	"github.com/jzes/CPPlayList/spotify"
	"github.com/jzes/CPPlayList/startmenu"
	"github.com/jzes/CPPlayList/state"
)

func main() {
	handlers := map[state.State]state.Handler{
		state.Menu:         startmenu.NewHandler(),
		state.LoginDeezer:  deezer.NewLoginHandler(),
		state.LoginSpotify: spotify.NewLoginHandler(),
	}

	p := tea.NewProgram(screen.New(handlers))
	if _, err := p.Run(); err != nil {
		fmt.Printf("deu ruim %+v", err)
		os.Exit(1)
	}
}
