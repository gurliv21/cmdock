package tui

import(
	tea "github.com/charmbracelet/bubbletea"
	"cmdock/internal/store"
)


func(m Model) Init() tea.Cmd{
	return nil
}

func Run() error{
	db, err:= store.InitDB()
	if err != nil{
		return err
	}
	defer db.Close()

	commands, err := store.ShowCommands(db)
	if err!=nil{
		return err
	}

	// m := Model{
	// 	Commands:commands,
	// }

	m :=NewModel()
	m.Commands = commands
	m.SearchCommands = commands

	p := tea.NewProgram(m, tea.WithAltScreen())
	_,err =p.Run()
	return err
}