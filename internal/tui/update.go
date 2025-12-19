package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
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
		case "q", "ctrl + c":
			return m, tea.Quit
		}
		switch m.step {

		case stepURL:
			if msg.Type == tea.KeyEnter {
				url := m.textInput.Value()
				m.step = stepLoading
				return m, tea.Batch(
					m.spinner.Tick,
					fetchFormatsCmd(url, m.app.GetFormats),
				)
			}

		case stepFormats:

			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)

			if msg.Type == tea.KeyEnter {
				m.step = stepDone
			}

			return m, cmd

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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

	return m, cmd
}

func (m *Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	switch m.step {

	case stepURL:
		return fmt.Sprintf(
			"Enter video URL:\n\n%s\n\nPress Enter to continue",
			m.textInput.View(),
		)

	case stepLoading:
		return m.spinner.View() + " Fetching formats\n"

	case stepFormats:
		return m.list.View()

	case stepDone:
		return "Downloading...\n"
	}

	return ""
}
