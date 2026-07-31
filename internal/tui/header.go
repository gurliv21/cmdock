package tui

import (
	"fmt"
	"strings"
	// "os"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) header() string {
    innerW := max(1, m.contentWidth()-headerStyle.GetHorizontalFrameSize())

	left :=titleStyle.Render("[workspace]")

	if m.ShowSearch {
	left = lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		m.Search.View(),
	)
}

    right := mutedStyle.Render(fmt.Sprintf("%d Total Commands", len(m.Commands)))

    gap := innerW - lipgloss.Width(left) - lipgloss.Width(right)
    content := left + strings.Repeat(" ", max(1, gap)) + right
	// content :=left

	// return fmt.Sprintf("%q", m.Search.Value())

    return headerStyle.
        Width(innerW).
        Render(content)
}