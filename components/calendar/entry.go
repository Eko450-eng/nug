package calendar

import (
	"nug/structs"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Selected int
	styles   structs.Styles
}

func InitModel() Model {
	return Model{
		Selected: time.Now().Day(),
		styles:   *structs.DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
