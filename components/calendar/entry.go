package calendar

import (
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Selected int
	styles   structs.Styles
}

func InitModel() Model {
	return Model{
		Selected: 0,
		styles:   *structs.DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
