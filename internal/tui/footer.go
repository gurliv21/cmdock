package tui

import(
	// "github.com/charmbracelet/lipgloss"
)

func (m Model) footer() string {
    frameW, _ := rootStyle.GetFrameSize()
    innerWidth := m.Width - frameW

    fFrameW, _ := footerStyle.GetFrameSize()

    return footerStyle.
        Width(innerWidth - fFrameW).
        Render("↑↓ Move    Enter Details  Esc Back   / Search    c Copy    d Delete    q Quit ")
}