package projectview

import (
	"nug/helpers"
	"nug/structs"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, structs.Keymap.QuitEasy) {
			m.Finished = true
			return m, nil
		}
	}

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	if m.Form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()
		for _, project := range m.projects {
			name := m.Form.GetString(strconv.Itoa(int(project.ID)))
			var p structs.Project
			p.ID = project.ID
			p.Name = name

			if name == "" {
				db.Delete(&p)
			} else {
				db.Save(&p)
			}

			m.Finished = true
		}
	}

	return m, cmd
}
