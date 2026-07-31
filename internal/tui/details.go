package tui 

import(
	"github.com/charmbracelet/lipgloss"
	"time"
	"fmt"

)

func (m Model) body() string{
	rootFrameH := rootStyle.GetVerticalFrameSize()
	headerH := lipgloss.Height(m.header())
	footerH := lipgloss.Height(m.footer())

	availH := m.Height - rootFrameH - headerH - footerH
	if availH < 0 {
		availH = 0
	}

	if !m.ShowDetails{
		return m.topPanel(availH)
	}
	topH := availH / 2
	bottomH := availH - topH

	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.topPanel(topH),
		m.bottomPanel(bottomH),
	)
}


func (m Model) topPanel(h int) string{
	visible := max(1,h)
	offset := 0
	if len(m.Commands) > visible {
		if m.Cursor >= visible {
			offset = m.Cursor - visible + 1
		}
		maxOffset := len(m.Commands) - visible
		if offset > maxOffset {
			offset = maxOffset
		}
	}

	end := offset + visible
	if end > len(m.Commands) {
		end = len(m.Commands)
	}

   var rows []string

   for i := offset;i<end;i++{
	cmd :=m.Commands[i]
	text := cmd.Command

	icon := "✓"
	if cmd.ExitCode !=0{
		icon ="✗"
	}

	row := icon + " " + text

	if i == m.Cursor{
		row = selectedStyle.Render(row)
	}else{
		row = normalStyle.Render(row)
	}

	rows = append(rows,row)

   }

   return lipgloss.NewStyle().
		Height(max(1,h)).
		MaxHeight(h).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				rows...,
			),
		)
}

func (m Model) bottomPanel(h int) string {
	var content string

	if m.Cursor >= 0 && m.Cursor < len(m.Commands) {
		cmd := m.Commands[m.Cursor]

		status := successStyle.Render("✓ success")
		if cmd.ExitCode != 0 {
			status = errorStyle.Render(fmt.Sprintf("✗ exit %d", cmd.ExitCode))
		}

		start := time.Unix(cmd.StartTime, 0)
		end := time.Unix(cmd.EndTime, 0)
		duration := end.Sub(start)

		rows := []string{
			mutedStyle.Render("Command  ") + normalStyle.Render(cmd.Command),
			mutedStyle.Render("Directory") + " " + normalStyle.Render(cmd.Directory),
			mutedStyle.Render("Status   ") + " " + status,
			mutedStyle.Render("Started  ") + " " + normalStyle.Render(start.Format("2006-01-02 15:04:05")),
			mutedStyle.Render("Duration ") + " " + normalStyle.Render(duration.String()),
		}

		content = lipgloss.JoinVertical(
			lipgloss.Left,
			rows...,
		)
	} else {
		content = mutedStyle.Render("No command selected")
	}

	return panelStyle.
		Height(max(1, h-panelStyle.GetVerticalFrameSize())).
		MaxHeight(h).
		Render(content)
}