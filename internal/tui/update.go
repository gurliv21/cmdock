package tui

import(
	 tea "github.com/charmbracelet/bubbletea"
)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "up":

            if m.Cursor > 0 {
                m.Cursor--
            }

        case "down":

            if m.Cursor < len(m.Commands)-1 {
                m.Cursor++
            }

        case "q", "ctrl+c":

            return m, tea.Quit
        }

    }

    return m, nil
}