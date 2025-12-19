package tui

import (
	"github.com/ayushsharma74/lazydlp/internal/app"
	"github.com/ayushsharma74/lazydlp/internal/domain"
	"github.com/charmbracelet/bubbles/list"
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
	app       *app.App
	step      step
	textInput textinput.Model
	formats   []domain.Format
	cursor    int
	spinner   spinner.Model
	list      list.Model
	width     int
	height    int

	listReady bool

	err error
}

func NewModel(a *app.App) *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.CharLimit = 150
	ti.Width = 20
	sp := spinner.New()
	sp.Spinner = spinner.Line

	return &Model{
		app:       a,
		textInput: ti,
		spinner:   sp,
		step:      stepURL,
	}
}
