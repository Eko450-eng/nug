package taskoverview

import (
	"fmt"
	"nug/components/calendar"
	"nug/components/createproject"
	"nug/components/createtask"
	"nug/components/settings"
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
	createProjectState
	calendarState
	settingState
)

type orderState int

const (
	prioAsc orderState = iota
	prioDesc
	none
)

type Model struct {
	Tasks         []structs.Task
	show_deleted  bool
	Cursor        int
	selected      map[int]struct{}
	styles        structs.Styles
	state         sessionState
	taskcard      taskcard.TaskCardModel
	createmodel   createtask.CreateModel
	ordering      orderState
	hideCompleted bool
	createProject createproject.Model
	calendar      calendar.Model
	settings      settings.Model

	Width  int
	Height int

	loaded bool
}

func (m Model) UpdateTasks() []structs.Task {
	var tasks []structs.Task
	db, _ := helpers.ConnectToSQLite()

	filtering := ""
	whereClause := ""

	switch m.ordering {
	case prioAsc:
		filtering = "prio ASC"
	case prioDesc:
		filtering = "prio DESC"
	default:
		filtering = ""
	}

	if !m.show_deleted {
		whereClause += "deleted = 0"
	}

	if m.hideCompleted && !m.show_deleted {
		whereClause += " AND completed = 0"
	} else if m.hideCompleted {
		whereClause += "completed = 0"
	}

	if res := db.
		Where(whereClause).
		Order(filtering).
		Find(&tasks); res.Error != nil {
		helpers.LogToFile(fmt.Sprintf("%e", res.Error))
	}
	return tasks
}

func InitModel() Model {
	var taskCard taskcard.TaskCardModel

	preSetSettings := helpers.GetSettings()
	ordering := preSetSettings.Ordering
	hideCompleted := false

	if preSetSettings.HideCompleted == 1 {
		hideCompleted = true
		helpers.LogToFile("True")
	}

	m := Model{
		selected:      make(map[int]struct{}),
		styles:        *structs.DefaultStyles(),
		show_deleted:  false,
		state:         mainState,
		taskcard:      taskCard,
		createmodel:   createtask.CreateModel{},
		createProject: createproject.InitModel(),
		settings:      settings.InitModel(),
		hideCompleted: hideCompleted,
		ordering:      orderState(ordering),
	}

	m.Tasks = m.UpdateTasks()

	if len(m.Tasks) > 0 {
		taskCard = taskcard.InitModel(m.Tasks[0], false)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
