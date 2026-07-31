package tui

import(
	 tea "github.com/charmbracelet/bubbletea"
)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.Width = msg.Width
        m.Height = msg.Height    

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
        case "esc":
            if m.ShowDetails{
                m.ShowDetails = false
            }
            if m.ShowSearch{
                m.ShowSearch = false
                m.Search.SetValue("")
                m.Search.Blur()
                m.Commands = m.Commands
                m.Cursor =0
            }   
        case "enter":
            m.ShowDetails = true
        
        case "/":
            m.ShowSearch = true 
            m.Search.Focus()  

        case "q", "ctrl+c":

            return m, tea.Quit
        }

    }
    var cmd tea.Cmd

    if m.ShowSearch{
        m.Search, cmd = m.Search.Update(msg)
        m.FilterCommands()
    }
    

    return m, cmd
}