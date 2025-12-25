package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ayushsharma74/lazydlp/internal/domain"
	"github.com/ayushsharma74/lazydlp/internal/util"
	"github.com/charmbracelet/bubbles/list"
)

type formatItem domain.Format

func (f formatItem) Title() string {
	if isAudioOnly(domain.Format(f)) {
		return f.Ext
	}
	return  videoHeight(f.Resolution)
}

func (f formatItem) Description() string {

	return fmt.Sprintf("%s • %s • %s", f.Resolution, f.Ext,
		util.HumanSize(f.Size))
}

func (f formatItem) FilterValue() string {
	return f.ID
}

func isAudioOnly(f domain.Format) bool {
	return strings.Contains(strings.ToLower(f.Resolution), "audio")
}

func videoHeight(res string) string {
	res = strings.ToLower(res)

	if strings.Contains(res, "audio") {
		return "audio only"
	}

	// Expect WIDTHxHEIGHT
	parts := strings.Split(res, "x")
	if len(parts) != 2 {
		return res // fallback, don't break UI
	}

	return parts[1] + "p"
}

func NewFormatList(formats []domain.Format) list.Model {
	sort.SliceStable(formats, func(i, j int) bool {
		fi := formats[i]
		fj := formats[j]

		fiAudio := isAudioOnly(fi)
		fjAudio := isAudioOnly(fj)

		// 1️⃣ Audio-only always last
		if fiAudio && !fjAudio {
			return false
		}
		if !fiAudio && fjAudio {
			return true
		}

		// 2️⃣ Both video → sort by height (ascending)
		if !fiAudio && !fjAudio {
			return videoHeight(fi.Resolution) < videoHeight(fj.Resolution)
		}

		// 3️⃣ Both audio → keep original order
		return false
	})

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
