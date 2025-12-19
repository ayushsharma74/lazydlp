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
