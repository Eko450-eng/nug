package taskcard

import (
	"nug/helpers"
	"nug/structs"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m TaskCardModel) UpdateTaskCard(msg tea.Msg) (TaskCardModel, tea.Cmd) {
	var cmd tea.Cmd
	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
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
		m.Exiting = true
		m.IsActive = false
	}

	return m, cmd
	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	switch m.Editing {
	// 	case true:
	// 		if key.Matches(msg, structs.Keymap.Quit) {
	// 			cmd = tea.Quit
	// 		} else {
	// 			// cmd = newCmd
	//
	// 			// return m, cmd
	// 			// if newState.Finished {
	// 			// 	m.state = 0
	// 			// 	m.Tasks = m.UpdateTasks()
	// 			// }
	// 		}
	// 	case false:
	// 		if key.Matches(msg, structs.Keymap.Edit) {
	// 			m.Editing = true
	// 		} else if key.Matches(msg, structs.Keymap.Up) && m.cursor > 0 {
	// 			m.cursor--
	// 		}
	// 		if key.Matches(msg, structs.Keymap.Down) && m.cursor < len(m.fields)-1 {
	// 			m.cursor++
	// 		}
	// 	}
	// }

	return m, cmd
}
