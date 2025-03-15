package settings

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

		var settings structs.Settings
		db.First(&settings)

		settings.HideCompleted = m.form.GetInt("hidecompleted")
		settings.Ordering = m.form.GetInt("ordering")

		db.Save(&settings)

		m.Finished = true
	}

	return m, cmd
}
