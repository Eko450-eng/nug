package createproject

import (
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	project  structs.Project
	Finished bool
	form     *huh.Form // huh.Form is just a tea.Model
}

func InitModel() Model {
	return Model{
		Finished: false,
		project:  structs.Project{},
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Key("name").
					Title("Name of the Project"),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}
