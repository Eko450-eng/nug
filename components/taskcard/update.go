package taskcard

import (
	"nug/helpers"
	"nug/structs"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m TaskCardModel) UpdateTaskCard(msg tea.Msg) (TaskCardModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.Editing {
		case true:
			if key.Matches(msg, structs.Keymap.Quit) {
				cmd = tea.Quit
			} else if key.Matches(msg, structs.Keymap.Save) {
				db, _ := helpers.ConnectToSQLite()
				db.Model(&structs.Task{}).
					Where("id = ?", m.Task.Id).
					Update(m.current.Question, helpers.NormalizeDate(m.current.InputField.Value())).
					Update("Updatedtime", time.Now())

				m.Editing = false
				m.Exiting = true
			} else {
				m.current.InputField, cmd = m.current.InputField.Update(msg)
			}
		case false:
			if key.Matches(msg, structs.Keymap.Edit) {
				m.Editing = true
				m.current = m.fields[m.cursor]
				switch m.fields[m.cursor].Question {
				case "Name":
					m.current.InputField.SetValue(m.Task.Name)
				case "Description":
					m.current.InputField.SetValue(m.Task.Description)
				case "Project_id":
					m.current.InputField.SetValue(strconv.Itoa(m.Task.Project_id))
				case "Prio":
					m.current.InputField.SetValue(strconv.Itoa(m.Task.Prio))
				case "Date":
					m.current.InputField.SetValue(helpers.NormalizeDate(m.Task.Date))
				}
			} else if key.Matches(msg, structs.Keymap.Save) {
				value := ""
				switch m.fields[m.cursor].Question {
				case "Name":
					value = m.Task.Name
				case "Description":
					value = m.Task.Description
				case "Project_id":
					value = strconv.Itoa(m.Task.Project_id)
				case "Prio":
					value = strconv.Itoa(m.Task.Prio)
				case "Date":
					value = helpers.NormalizeDate(m.Task.Date)
				}
				m.current.InputField.SetValue(value)
			} else if key.Matches(msg, structs.Keymap.Up) && m.cursor > 0 {
				m.cursor--
				m.current = m.fields[m.cursor]
			}
			if key.Matches(msg, structs.Keymap.Down) && m.cursor < len(m.fields)-1 {
				m.cursor++
				m.current = m.fields[m.cursor]
			}
		}
	}

	return m, cmd
}
