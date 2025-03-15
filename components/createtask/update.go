package createtask

import (
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m CreateModel) UpdateCreateElement(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd
	// m.current = m.Fields[m.EditLine]

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()

		m.Newtask = structs.Task{
			Name:        m.form.GetString("name"),
			Description: m.form.GetString("description"),
			Prio:        m.form.GetInt("prio"),
			Project_id:  m.form.GetInt("project"),
			Deleted:     0,
		}

		db.Create(&m.Newtask)

		m.Exiting = true
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
			return m, cmd
		}
	}

	return m, cmd
}
