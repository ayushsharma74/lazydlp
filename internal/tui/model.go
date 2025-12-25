package tui

import (
	"github.com/ayushsharma74/lazydlp/internal/app"
	"github.com/ayushsharma74/lazydlp/internal/domain"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
)

type step int

const (
	stepURL step = iota
	stepLoading
	stepFormats
	stepDone
)

type Model struct {
	app         *app.App
	step        step
	textInput   textinput.Model
	formats     []domain.Format
	spinner     spinner.Model
	url         string
	list        list.Model
	width       int
	height      int
	progress    progress.Model
	downloadSpeed string
	downloadErr error
	listReady   bool

	err error
}

func NewModel(a *app.App) *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.CharLimit = 150
	ti.Width = 50
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	p := progress.New(progress.WithDefaultGradient())
	p.ShowPercentage = false

	return &Model{
		app:       a,
		textInput: ti,
		spinner:   sp,
		step:      stepURL,
		progress:  p,
	}
}
