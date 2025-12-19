package main

import (
	"log"
	"os"

	"github.com/ayushsharma74/lazydlp/internal/app"
	"github.com/ayushsharma74/lazydlp/internal/config"
	"github.com/ayushsharma74/lazydlp/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := os.OpenFile("lazydlp.debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := config.Load()

	a := app.New(cfg)

	p := tea.NewProgram(
		tui.NewModel(a),
		tea.WithAltScreen(),
	)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
