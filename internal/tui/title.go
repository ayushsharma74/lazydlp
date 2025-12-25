package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7C7CFF")).
		Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
		Faint(true).
		Foreground(lipgloss.Color("#A1A1AA")).
		Align(lipgloss.Center)

	containerStyle = lipgloss.NewStyle().
		Padding(1, 2)
)

func titleArt() string {
	return (`
 ██▓    ▄▄▄      ▒███████▒▓██   ██▓▓█████▄  ██▓     ██▓███     
▓██▒   ▒████▄    ▒ ▒ ▒ ▄▀░ ▒██  ██▒▒██▀ ██▌▓██▒    ▓██░  ██▒   
▒██░   ▒██  ▀█▄  ░ ▒ ▄▀▒░   ▒██ ██░░██   █▌▒██░    ▓██░ ██▓▒   
▒██░   ░██▄▄▄▄██   ▄▀▒   ░  ░ ▐██▓░░▓█▄   ▌▒██░    ▒██▄█▓▒ ▒   
░██████▒▓█   ▓██▒▒███████▒  ░ ██▒▓░░▒████▓ ░██████▒▒██▒ ░  ░   
░ ▒░▓  ░▒▒   ▓▒█░░▒▒ ▓░▒░▒   ██▒▒▒  ▒▒▓  ▒ ░ ▒░▓  ░▒▓▒░ ░  ░   
░ ░ ▒  ░ ▒   ▒▒ ░░░▒ ▒ ░ ▒ ▓██ ░▒░  ░ ▒  ▒ ░ ░ ▒  ░░▒ ░        
  ░ ░    ░   ▒   ░ ░ ░ ░ ░ ▒ ▒ ░░   ░ ░  ░   ░ ░   ░░          
    ░  ░     ░  ░  ░ ░     ░ ░        ░        ░  ░            
                 ░         ░ ░      ░                             
	`)
}

func subtitle() string {
	return "minimal downloader"
}

func StyledTitle(width int) string {
	art := titleStyle.
		Width(width).
		Render(titleArt())

	sub := subtitleStyle.
		Width(width).
		Render(subtitle())

	return containerStyle.
		Width(width).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				art,
				"",
				sub,
			),
		)
}
