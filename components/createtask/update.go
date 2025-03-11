package createtask

import (
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m CreateModel) UpdateCreateElement(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd

	m.current = m.Fields[m.EditLine]
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
			return m, cmd
		} else if key.Matches(msg, structs.Keymap.Save) {
			if m.EditLine < len(m.Fields)-1 {
				switch m.current.Question {
				case "Name":
					m.Newtask.Name = m.current.InputField.Value()
				case "Description":
					m.Newtask.Description = m.current.InputField.Value()
				case "Project_id":
					m.Newtask.Project_id = helpers.SetDefaultInt(m.current.InputField.Value())
				case "Prio":
					m.Newtask.Prio = helpers.SetDefaultInt(m.current.InputField.Value())
				case "Completed":
					m.Newtask.Completed = helpers.SetDefaultInt(m.current.InputField.Value())
				case "Date":
					m.Newtask.Date = helpers.NormalizeDate(m.createInputFields.Value())
				}
				m.EditLine++
			} else {
				m.EditLine = 0
				db, _ := helpers.ConnectToSQLite()

				m.Newtask.Date = helpers.NormalizeDate(m.Newtask.Date)

				db.Create(&m.Newtask)
				m.Newtask = helpers.Resettask()
				m.Exiting = true
			}
		} else {
			m.current.InputField, cmd = m.current.InputField.Update(msg)
		}
	}

	return m, cmd
}
