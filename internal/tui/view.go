package tui

import (
	"fmt"
)

func (m Model) View() string {

    s := "Recent Commands\n\n"

    for i, cmd := range m.Commands {

        cursor := " "

        if i == m.Cursor {
            cursor = ">>"
        }

        s += fmt.Sprintf("%s %s\n", cursor, cmd.Command)
    }

    s += "\n\n↑ ↓ Move   q Quit"

    return s
}