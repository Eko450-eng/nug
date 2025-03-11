package taskoverview

import (
	"fmt"
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

	Width  int
	Height int
}

func (m Model) UpdateTasks() []structs.Task {
	var tasks []structs.Task
	db, _ := helpers.ConnectToSQLite()
	if m.show_deleted {
		if res := db.
			Find(&tasks); res.Error != nil {
			helpers.LogToFile(fmt.Sprintf("%e", res.Error))
		}
		return tasks
	} else {
		if res := db.
			Where("deleted = ?", 0).
			Find(&tasks); res.Error != nil {
			helpers.LogToFile(fmt.Sprintf("%e", res.Error))
		}
		return tasks
	}
}

func InitModel() Model {
	var taskCard taskcard.TaskCardModel

	tasks := helpers.UpdateTasks()

	if len(tasks) > 0 {
		taskCard = taskcard.InitModel(tasks[0], false)
	}

	return Model{
		selected:     make(map[int]struct{}),
		Tasks:        tasks,
		styles:       *structs.DefaultStyles(),
		show_deleted: false,
		state:        mainState,
		taskcard:     taskCard,
		createmodel:  createtask.CreateModel{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
