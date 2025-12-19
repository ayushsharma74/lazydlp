package app

import (
	"github.com/ayushsharma74/lazydlp/internal/config"
	"github.com/ayushsharma74/lazydlp/internal/ytdlp"
)

type App struct {
	Yt *ytdlp.Client
}

func New(cfg config.Config) *App {
	return &App{
		Yt: ytdlp.New(cfg.YtDlpPath),
	}
}