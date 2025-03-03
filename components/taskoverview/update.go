package taskoverview

import (
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Up) && m.Cursor > 0 {
			m.Cursor--
		} else if key.Matches(msg, structs.Keymap.Down) && m.Cursor < len(m.Tasks)-1 {
			m.Cursor++
		} else if key.Matches(msg, structs.Keymap.ShowDeleted) {
			m.Cursor = 0
			m.Show_deleted = !m.Show_deleted
			m.Tasks = helpers.UpdateTasks(m.Show_deleted)
		} else if key.Matches(msg, structs.Keymap.Delete) {
			db, _ := helpers.ConnectToSQLite()
			currenttask := m.Tasks[m.Cursor]
			if currenttask.Deleted == 1 {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.Id).
					Update("Deleted", 0)
				m.Tasks = helpers.UpdateTasks(false)
				m.Show_deleted = false
				m.Cursor = 0
			} else {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.Id).
					Update("Deleted", 1)
				m.Tasks = helpers.UpdateTasks(m.Show_deleted)
				if m.Cursor == 0 {
					m.Cursor = 0
				} else {
					m.Cursor--
				}
			}
		} else if key.Matches(msg, structs.Keymap.Check) {
			db, _ := helpers.ConnectToSQLite()
			currenttask := m.Tasks[m.Cursor]
			newvalue := 0
			if currenttask.Completed == 1 {
				newvalue = 0
			} else {
				newvalue = 1
			}
			db.Model(&structs.Task{}).
				Where("id = ?", currenttask.Id).
				Update("Completed", newvalue)
			m.Tasks = helpers.UpdateTasks(m.Show_deleted)
		}
	}
	return m, cmd
}
