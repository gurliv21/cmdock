package tui

import (
	// "fmt"
    // "strings"
    "github.com/charmbracelet/lipgloss"
)

func (m Model) View() string{
    root :=rootStyle.
    Width(m.contentWidth())

    return root.Render(
        lipgloss.JoinVertical(
        lipgloss.Left,
        m.header(),
        m.body(),
        m.footer(),
    ),
    )
}

func (m Model) contentWidth() int {
	w := m.Width - rootStyle.GetHorizontalFrameSize()
	if w < 1 {
		w = 1
	}
	return w
}