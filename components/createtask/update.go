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

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	if m.Form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()

		m.Newtask = structs.Task{
			Name:        m.Form.GetString("name"),
			Description: m.Form.GetString("description"),
			Prio:        m.Form.GetInt("prio"),
			Project_id:  m.Form.GetInt("project"),
			Deleted:     0,
		}

		db.Create(&m.Newtask)

		m.Finished = true
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
			return m, cmd
		} else if key.Matches(msg, structs.Keymap.QuitEasy) {
			m.Finished = true
			return m, cmd
		}
	}

	return m, cmd
}
