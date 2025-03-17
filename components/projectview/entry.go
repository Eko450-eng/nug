package projectview

import (
	"fmt"
	"nug/helpers"
	"nug/structs"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	projects []structs.Project
	Finished bool
	Form     *huh.Form
}

func formTemplate(name string, id int) huh.Field {
	taskCount := helpers.GetTaskCountOfProject(uint(id))

	formTemplate := huh.NewInput().
		Title(fmt.Sprintf("%s - Tasks: %d", name, taskCount)).
		Key(strconv.Itoa(id)).
		Placeholder("Rename...").
		Value(&name)

	return formTemplate
}

func InitModel() Model {
	projects := helpers.GetProjects()
	var formFields []huh.Field

	for _, project := range projects {
		formFields = append(formFields, formTemplate(project.Name, int(project.ID)))
	}

	return Model{
		Finished: false,
		projects: projects,
		Form: huh.NewForm(
			huh.NewGroup(
				formFields...,
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Form.Init()
}
