package tui

import(
	"cmdock/internal/store"
)

type Model struct{
	Commands []store.Command
	Cursor  int
}