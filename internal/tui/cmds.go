package tui

import (
	"github.com/ayushsharma74/lazydlp/internal/domain"
	tea "github.com/charmbracelet/bubbletea"
)

func fetchFormatsCmd(url string, appFn func(string) ([]domain.Format, error)) tea.Cmd {
	return func() tea.Msg {
		formats, err := appFn(url)
		if err != nil {
			return errMsg{err}
		}
		return formatsMsg{formats}
	}
}

func downloadCmd(url string, formatID string, downloadFn func(string, string, chan<- domain.ProgressUpdate) error) tea.Cmd {
	return func() tea.Msg {
		// Updated channel type and buffer
		progressCh := make(chan domain.ProgressUpdate, 100)

		go func() {
			_ = downloadFn(url, formatID, progressCh)
			close(progressCh)
		}()

		return beginDownloadMsg{ch: progressCh}
	}
}

func waitForProgress(ch chan domain.ProgressUpdate) tea.Cmd {
	return func() tea.Msg {
		update, ok := <-ch
		if !ok {
			return downloadFinishedMsg{err: nil}
		}
		return downloadProgressMsg{update: update, ch: ch}
	}
}
