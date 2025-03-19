package taskcard

import (
	"nug/helpers"
	"nug/structs"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m TaskCardModel) UpdateTaskCard(msg tea.Msg) (TaskCardModel, tea.Cmd) {
	var cmd tea.Cmd
	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, structs.Keymap.QuitEasy) {
			m.Editing = false
			m.Finished = true
			m.IsActive = false
			return m, nil
		}
	}

	if m.Form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()

		newTask := structs.Task{
			Name:        m.Form.GetString("name"),
			Description: m.Form.GetString("description"),
			Project_id:  m.Form.GetInt("project"),
			Prio:        m.Form.GetInt("prio"),
			Date:        helpers.NormalizeDate(m.Form.GetString("date")),
			Modified:    time.Now().String(),
		}

		db.Model(&m.Task).
			Updates(newTask)

		m.Editing = false
		m.Finished = true
		m.IsActive = false
	}

	return m, cmd
}
