package taskoverview

import (
	"fmt"
	"nug/helpers"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View(width, height int) string {
	tasks := ""

	if m.createmodel.Exiting {
		m.state = mainState
		m.Tasks = m.UpdateTasks()
		m.createmodel.Exiting = false
	}

	if m.taskcard.Exiting {
		m.Tasks = m.UpdateTasks()
		m.state = mainState

		m.taskcard.Task = m.Tasks[m.Cursor]

		m.taskcard.IsActive = false
		m.taskcard.Exiting = false
	}

	m.createProject.Init()

	if m.createProject.Finished {
		m.state = mainState
		m.createProject.Finished = false
	}

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColor)
	borderStyleActive := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColorActive)
	tasksBox := []string{}
	tasksBox = append(tasksBox, lipgloss.JoinHorizontal(
		lipgloss.Top,
		"Prio",
		"  ",
		"Project",
		" - ",
		"Task",
	))

	for i, task := range m.Tasks {

		checked := lipgloss.NewStyle().
			Background(lipgloss.Color("#e9c46a")).
			Foreground(lipgloss.Color("#000000"))

		if m.Tasks[i].Completed == 1 {
			checked = checked.
				Background(lipgloss.Color("#a7c957")).
				Foreground(lipgloss.Color("#000000"))
		}

		cursor := "  "
		if m.Cursor == i {
			checked = checked.Bold(true)
			cursor = "> "
		}

		if m.Tasks[i].Deleted == 1 {
			checked = checked.
				Background(lipgloss.Color("#780000")).
				Foreground(lipgloss.Color("#ffffff"))
		}

		taskText := fmt.Sprintf("%s%s - %s", cursor, helpers.GetProjectName(task.Project_id), helpers.ShortenString(task.Name, 30))
		tasksBox = append(tasksBox, lipgloss.JoinHorizontal(
			lipgloss.Top,
			strconv.Itoa(task.Prio),
			"  ",
			checked.Render(taskText),
		))
	}
	tasks = lipgloss.JoinVertical(
		lipgloss.Top,
		tasksBox...,
	)

	width = width - 10
	height = height - 2

	leftWidth := width / 3
	rightWidth := width - leftWidth

	res := ""

	switch m.state {
	case settingState:
		res += lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			m.settings.View(),
		)
	case calendarState:
		res += lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			m.calendar.View(width, height),
		)
	case mainState:
		res += lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyleActive.Width(leftWidth).Height(height).Render(
					tasks,
				),
				borderStyle.Width(rightWidth).Height(height).Render(
					m.taskcard.View(rightWidth-30),
				),
			),
		)

	case createProjectState:
		res = m.createProject.View()

	case infoState:
		res += lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyle.Width(leftWidth).Height(height).Render(
					tasks,
				),
				borderStyleActive.Width(rightWidth).Height(height).Render(
					m.taskcard.View(rightWidth-30),
				),
			),
		)
	case createState:
		// if len(m.createmodel.Fields) == 0 {
		// 	return "No fields to display."
		// } else {

		// current := m.createmodel.Fields[m.createmodel.EditLine]

		m.createmodel.Init()

		res += lipgloss.Place(
			width,
			height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				// current.Question,
				m.createmodel.View(),
			),
		)
		// }
	}

	return res
}
