package createtask

import (
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type CreateModel struct {
	EditLine int
	Fields   []structs.Questions
	Newtask  structs.Task
	styles   structs.Styles
	Finished bool
	Form     *huh.Form
}

func InitTaskCreation() CreateModel {
	projectItems := helpers.GetProjectItems()

	return CreateModel{
		EditLine: 0,
		Newtask:  structs.Task{},
		styles:   *structs.DefaultStyles(),
		Finished: false,
		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Key("name").
					Title("Name"),
				huh.NewText().
					Key("description").
					Title("Description"),
				huh.NewSelect[int]().
					Key("prio").
					Title("Priority").
					Options(
						huh.NewOption("1", 1),
						huh.NewOption("2", 2),
						huh.NewOption("3", 3),
					),
				huh.NewSelect[int]().
					Key("project").
					Title("Project").
					Options(
						projectItems...,
					),
			),
		),
	}
}

func (m CreateModel) Init() tea.Cmd {
	return m.Form.Init()
}
