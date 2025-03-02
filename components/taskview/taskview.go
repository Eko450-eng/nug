package taskview

import (
	"fmt"
	"nask/helpers"
	"nask/inputs"
	"nask/structs"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Fields struct {
	Question   string
	Answer     string
	InputField inputs.Input
}

type TaskViewModel struct {
	Task        structs.Task
	Exiting     bool
	styles      structs.Styles
	editing     bool
	editline    int
	cursor      int
	fields      []Fields
	current     Fields
	ChangeEvent bool
	isActive    bool
}

func newQuestion(q string) Fields {
	return Fields{Question: q}
}

func newShortQuestion(q string) Fields {
	question := newQuestion(q)
	model := inputs.NewShortAnswerField()
	question.InputField = model
	return question
}

func newLongQuestion(q string) Fields {
	question := newQuestion(q)
	model := inputs.NewLongAnswerField()
	question.InputField = model
	return question
}

func InitTaskCreation(task structs.Task, isActive bool) TaskViewModel {
	fields := []Fields{
		newShortQuestion("Name"),
		newLongQuestion("Description"),
		newShortQuestion("Project_id"),
		newShortQuestion("Prio"),
	}

	return TaskViewModel{
		Task:        task,
		cursor:      0,
		styles:      *structs.DefaultStyles(),
		fields:      fields,
		current:     fields[0],
		Exiting:     false,
		ChangeEvent: false,
		isActive:    isActive,
	}
}

func (m TaskViewModel) Update(msg tea.Msg) (TaskViewModel, tea.Cmd) {
	var cmd tea.Cmd
	switch m.editing {
	case true:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c":
				cmd = tea.Quit
				return m, cmd
			case "enter", "l":
				db, _ := helpers.ConnectToSQLite()
				db.Model(&structs.Task{}).
					Where("id = ?", m.Task.Id).
					Update(m.current.Question, m.current.InputField.Value())
				m.Exiting = true
				m.editing = false
			default:
				m.current.InputField, cmd = m.current.InputField.Update(msg)
			}
		}
	case false:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c":
				cmd = tea.Quit
			case "q", "h":
				m.Exiting = true
			case "enter":
				m.editing = true
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
				}
				m.current.InputField.SetValue(value)
			case "k", "up":
				if m.cursor > 0 {
					m.cursor--
					m.current = m.fields[m.cursor]
				}
			case "j", "down":
				if m.cursor < len(m.fields)-1 {
					m.cursor++
					m.current = m.fields[m.cursor]
				}
			}
		}
	}

	return m, cmd
}

func selectedCursor(cursor, pos int, isActive bool) string {
	c := ""
	if cursor == pos && isActive {
		c = ">"
	}
	return c
}

func (m TaskViewModel) View(width int) string {
	borderColor := m.styles.BorderColor
	if m.Task.Deleted == 1 {
		borderColor = lipgloss.Color("9")
	}
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Width(width)

	if m.editing {
		return m.styles.InputField.Render(m.current.InputField.View())
	} else {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			borderStyle.Render(
				"Name:\n",
				fmt.Sprintf("%s %s", selectedCursor(m.cursor, 0, m.isActive), m.Task.Name),
			),
			borderStyle.Render(
				"Description:\n",
				fmt.Sprintf("%s %s", selectedCursor(m.cursor, 1, m.isActive), m.Task.Description),
			),
			borderStyle.Render(
				"Project_id:\n",
				fmt.Sprintf("%s %s", selectedCursor(m.cursor, 2, m.isActive), strconv.Itoa(m.Task.Project_id)),
			),
			borderStyle.Render(
				"Prio:\n",
				fmt.Sprintf("%s %s", selectedCursor(m.cursor, 3, m.isActive), strconv.Itoa(m.Task.Prio)),
			),
		)
	}
}
