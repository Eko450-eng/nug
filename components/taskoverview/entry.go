package taskoverview

import (
	"nug/components/createtask"
	"nug/components/taskcard"
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	mainState sessionState = iota
	infoState
	createState
)

type Model struct {
	Tasks        []structs.Task
	show_deleted bool
	Cursor       int
	selected     map[int]struct{}
	styles       structs.Styles
	state        sessionState
	taskcard     taskcard.TaskCardModel
	createmodel  createtask.CreateModel
}

func (m Model) UpdateTasks() []structs.Task {
	var tasks []structs.Task
	db, _ := helpers.ConnectToSQLite()
	if m.show_deleted {
		if res := db.
			Find(&tasks); res.Error != nil {
			panic(res.Error)
		}
		return tasks
	} else {
		if res := db.
			Where("deleted = ?", 0).
			Find(&tasks); res.Error != nil {
			panic(res.Error)
		}
		return tasks
	}
}

func InitModel() Model {

	tasks := helpers.UpdateTasks()

	return Model{
		selected:     make(map[int]struct{}),
		Tasks:        tasks,
		styles:       *structs.DefaultStyles(),
		show_deleted: false,
		state:        mainState,
		taskcard:     taskcard.InitModel(tasks[0], false),
		createmodel:  createtask.CreateModel{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
