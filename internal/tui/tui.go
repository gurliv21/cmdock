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

	m := Model{
		Commands:commands,
	}
	_,err = tea.NewProgram(m).Run()
	return err
}