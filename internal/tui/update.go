package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.textInput.Cursor.BlinkCmd(),
		m.spinner.Tick,
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

		switch m.step {

		case stepURL:
			if msg.Type == tea.KeyEnter {
				url := m.textInput.Value()
				m.url = url
				m.step = stepLoading
				return m, tea.Batch(
					m.spinner.Tick,
					fetchFormatsCmd(url, m.app.GetFormats),
				)
			}

		case stepFormats:
			if msg.Type == tea.KeyEnter {
				if selected, ok := m.list.SelectedItem().(formatItem); ok {
					m.step = stepDone
					log.Print(selected.ID)
					formatID := selected.ID
					return m, downloadCmd(m.url, formatID, m.app.DownloadFormat)
				}
			}
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	case beginDownloadMsg:
		return m, waitForProgress(msg.ch)

	case downloadProgressMsg:
		log.Printf("TUI RECEIVED: %s", msg.update.Speed)
		m.downloadSpeed = msg.update.Speed

		// 1. Trigger the animation command using the actual percentage
		progressCmd := m.progress.SetPercent(msg.update.Percent)

		// 2. Update the progress model instance itself
		// This is the "missing link" that processes the animation logic
		newModel, cmd := m.progress.Update(msg)
		if p, ok := newModel.(progress.Model); ok {
			m.progress = p
		}

		// 3. Batch the internal progress command and our listener command
		return m, tea.Batch(progressCmd, cmd, waitForProgress(msg.ch))

	case downloadFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = msg.Width - padding*2 - 4

		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}

		if m.listReady {
			m.list.SetSize(msg.Width, msg.Height-4)
		}

	case spinner.TickMsg:
		if m.step == stepLoading {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case formatsMsg:
		m.formats = msg.formats
		m.step = stepFormats
		m.list = NewFormatList(msg.formats)
		m.listReady = true

		// APPLY SIZE IMMEDIATELY
		if m.width > 0 && m.height > 0 {
			m.list.SetSize(m.width, m.height-4)
		}

		return m, nil

	case errMsg:
		m.err = msg.err
		return m, nil
	}

	if m.step == stepURL {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	if m.step == stepDone {
		var progressCmd tea.Cmd
		var newModel tea.Model

		newModel, progressCmd = m.progress.Update(msg)
		if p, ok := newModel.(progress.Model); ok {
			m.progress = p
		}
		// Combine this with existing commands
		cmd = tea.Batch(cmd, progressCmd)
	}

	return m, cmd
}

func (m *Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	switch m.step {

	case stepURL:
		return fmt.Sprintf(
			"\n\n%s\n\n Enter video URL:\n\n%s\n\nPress Enter to continue",
			StyledTitle(m.width),
			m.textInput.View(),
		)

	case stepLoading:
		return fmt.Sprintf(
			"\n\n%s\n\n  %s Fetching Formats",
			StyledTitle(m.width),
			m.spinner.View(),
		)

	case stepFormats:
		return m.list.View()

	case stepDone:
		progressDone := m.progress.Percent()
		speed := m.downloadSpeed
		if progressDone == 0 {
			progressDone = 0.0
		}
		if speed == "" {
			speed = "calculating..."
		}
		return fmt.Sprintf(
			"\n\n%s\n\n%s\n\n%s %0.1f%%\n",
			StyledTitle(m.width),
			speed,
			m.progress.View(),
			progressDone*100,
		)
	}

	return ""
}
