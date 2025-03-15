package createproject

import (
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()

		m.project = structs.Project{
			Name:    m.form.GetString("name"),
			Deleted: 0,
		}

		db.Create(&m.project)

		m.Finished = true
	}

	return m, cmd
}
