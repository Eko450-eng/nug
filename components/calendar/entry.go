package calendar

import (
	"nug/structs"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Selected int
	styles   structs.Styles
	width    int
	height   int
}

func InitModel(width, height int) Model {
	return Model{
		Selected: time.Now().Day(),
		styles:   *structs.DefaultStyles(),

		width:  width,
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
