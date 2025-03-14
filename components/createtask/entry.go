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
	Exiting  bool
	form     *huh.Form
}

func InitTaskCreation() CreateModel {
	projectItems := helpers.GetProjectItems()

	return CreateModel{
		EditLine: 0,
		Newtask:  structs.Task{},
		styles:   *structs.DefaultStyles(),
		Exiting:  false,
		form: huh.NewForm(
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
						huh.NewOption("1", 0),
						huh.NewOption("2", 1),
						huh.NewOption("3", 2),
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
	return m.form.Init()
}
