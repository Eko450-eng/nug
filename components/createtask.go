package components

import (
	"nask/helpers"
	"nask/inputs"
	"nask/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateModel struct {
	EditLine          int
	Fields            []structs.Questions
	Newtask           structs.Task
	styles            structs.Styles
	Exiting           bool
	createInputFields textinput.Model
}

func newQuestion(q string) structs.Questions {
	return structs.Questions{Question: q}
}

func newShortQuestion(q string) structs.Questions {
	question := newQuestion(q)
	model := inputs.NewShortAnswerField()
	question.InputField = model
	return question
}

func newLongQuestion(q string) structs.Questions {
	question := newQuestion(q)
	model := inputs.NewLongAnswerField()
	question.InputField = model
	return question
}

func InitTaskCreation() CreateModel {
	questions := []structs.Questions{
		newShortQuestion("Name"),
		newLongQuestion("Description"),
		newShortQuestion("Project_id"),
		newShortQuestion("Prio"),
		newShortQuestion("Completed"),
	}

	return CreateModel{
		EditLine: 0,
		Fields:   questions,
		Newtask:  structs.Task{},
		styles:   *structs.DefaultStyles(),
		Exiting:  false,
	}
}

func (m CreateModel) UpdateCreateElement(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.Fields[m.EditLine]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			cmd = tea.Quit
			return m, cmd
		case "enter":
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
				}
				m.EditLine++
			} else {
				m.EditLine = 0
				db, _ := helpers.ConnectToSQLite()
				db.Create(&m.Newtask)
				m.Newtask = helpers.Resettask()
				m.Exiting = true
			}
		default:
			current.InputField, cmd = current.InputField.Update(msg)
			return m, cmd
		}
	}

	current.InputField, cmd = current.InputField.Update(msg)
	return m, cmd
}

func (m CreateModel) View() string {
	current := m.Fields[m.EditLine]
	return m.styles.InputField.Render(current.InputField.View())
}
