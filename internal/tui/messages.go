package tui

import "github.com/ayushsharma74/lazydlp/internal/domain"

type formatsMsg struct {
	formats []domain.Format
}

type errMsg struct {
	err error
}
