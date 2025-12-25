package tui

import "github.com/ayushsharma74/lazydlp/internal/domain"

type formatsMsg struct {
	formats []domain.Format
}

type errMsg struct {
	err error
}



type beginDownloadMsg struct {
	ch chan domain.ProgressUpdate
}

type downloadProgressMsg struct {
	update domain.ProgressUpdate
	ch     chan domain.ProgressUpdate
}

type downloadFinishedMsg struct {
	err error
}


