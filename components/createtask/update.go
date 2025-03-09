package createtask

import (
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m CreateModel) UpdateCreateElement(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.Fields[m.EditLine]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
			return m, cmd
		} else if key.Matches(msg, structs.Keymap.Save) {
			if m.EditLine < len(m.Fields)-1 {
				switch current.Question {
				case "Name":
					m.Newtask.Name = current.InputField.Value()
				case "Description":
					m.Newtask.Description = current.InputField.Value()
				case "Project_id":
					m.Newtask.Project_id = helpers.SetDefaultInt(current.InputField.Value())
				case "Prio":
					m.Newtask.Prio = helpers.SetDefaultInt(current.InputField.Value())
				case "Completed":
					m.Newtask.Completed = helpers.SetDefaultInt(current.InputField.Value())
				case "Time":
					m.Newtask.Time = current.InputField.Value()
				}
				m.EditLine++
			} else {
				m.EditLine = 0
				db, _ := helpers.ConnectToSQLite()
				db.Create(&m.Newtask)
				m.Newtask = helpers.Resettask()
				m.Exiting = true
			}
		} else {
			current.InputField, cmd = current.InputField.Update(msg)
		}
	}

	return m, cmd
}
