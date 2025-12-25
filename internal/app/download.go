package app

import "github.com/ayushsharma74/lazydlp/internal/domain"

func (a *App) GetFormats(url string) ([]domain.Format, error) {
	return a.Yt.ListFormats(url)
}

func (a *App) DownloadFormat(url string, formatID string, progressCh chan <- domain.ProgressUpdate) error {
    return a.Yt.Download(url, formatID, "~/Downloads", progressCh)
}
