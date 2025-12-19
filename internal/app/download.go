package app

import "github.com/ayushsharma74/lazydlp/internal/domain"

func (a *App) GetFormats(url string) ([]domain.Format, error) {
	return a.Yt.ListFormats(url)
}
