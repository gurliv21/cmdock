package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#89B4FA"))

	headerStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), false, false, true, false)

	panelStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#45475A")).
		Padding(1)

	selectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1E1E2E")).
		Background(lipgloss.Color("#89B4FA")).
		Bold(true)

	normalStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CDD6F4"))

	mutedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C7086"))

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A6E3A1"))

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F38BA8"))

	footerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C7086")).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Padding(0, 1)

	rootStyle = lipgloss.NewStyle().
    	Border(lipgloss.RoundedBorder()).
		Padding(1)

	sectionTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#89B4FA")).
		Padding(0, 1)

	sectionStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#45475A")).
		Padding(0, 1)

	contentStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CDD6F4"))	
)