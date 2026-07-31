package tui

import(
	"cmdock/internal/store"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct{
	Commands []store.Command
	SearchCommands []store.Command
	Cursor  int
	Width int
	Height int
	ShowDetails bool
	ShowSearch bool
	Search textinput.Model
}


func NewModel() Model{
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit =100
	ti.Width =40

	return Model{
		Search:ti,
	}
}