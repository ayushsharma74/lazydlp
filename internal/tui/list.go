package tui

import (
	"fmt"

	"github.com/ayushsharma74/lazydlp/internal/domain"
	"github.com/ayushsharma74/lazydlp/internal/util"
	"github.com/charmbracelet/bubbles/list"
)

type formatItem domain.Format

func (f formatItem) Title() string {
	return fmt.Sprintf("%s", f.Ext)
}

func (f formatItem) Description() string {

	return fmt.Sprintf("%s • %s • %s", f.Resolution, f.Ext,
		util.HumanSize(f.Size))
}

func (f formatItem) FilterValue() string {
	return f.ID
}

func NewFormatList(formats []domain.Format) list.Model {
	items := make([]list.Item, len(formats))
	for i, f := range formats {
		items[i] = formatItem(f)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select Format"
	l.SetFilteringEnabled(true)
	l.SetShowStatusBar(false)
	l.SetShowHelp(true)

	return l
}
