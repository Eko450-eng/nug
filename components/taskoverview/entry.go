package taskoverview

import (
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Tasks        []structs.Task
	Show_deleted bool
	Cursor       int
	Selected     map[int]struct{}
	styles       structs.Styles
}

func InitModel(Show_deleted bool) Model {
	return Model{
		Cursor:       0,
		Selected:     make(map[int]struct{}),
		Tasks:        helpers.UpdateTasks(Show_deleted),
		styles:       *structs.DefaultStyles(),
		Show_deleted: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
