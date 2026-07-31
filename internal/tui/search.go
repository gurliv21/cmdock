package tui

import(
   "cmdock/internal/store"
    "github.com/sahilm/fuzzy"	
	// "cmdock/internal/store"
)

func(m *Model) FilterCommands(){
	query :=m.Search.Value()
	if query == ""{
		m.Commands = m.SearchCommands
		return 
	}

	source := commandSource(m.SearchCommands)

	matches := fuzzy.FindFrom(query, source)

	filtered := make([]store.Command,0,len(matches))

	for _,match := range matches{
		filtered = append(filtered, m.SearchCommands[match.Index])
	}

	m.Commands = filtered

	if m.Cursor >= len(m.Commands) {
        m.Cursor = 0
    }
}


type commandSource []store.Command

func (c commandSource) String(i int) string {
	return c[i].Command
}

func (c commandSource) Len() int {
	return len(c)
}