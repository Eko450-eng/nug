package taskoverview

import (
	"fmt"

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

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColor)
	borderStyleActive := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColorActive)

	for i, task := range m.Tasks {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.Tasks[i].Completed == 1 {
			checked = "x"
		}
		if m.Tasks[i].Deleted == 1 {
			tasks += fmt.Sprintf("%s D-[%s] %s\n", cursor, checked, task.Name)
		} else {
			tasks += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Name)
		}
	}

	width = width - 10
	height = height - 2

	leftWidth := width / 3
	rightWidth := width - leftWidth

	res := ""

	switch m.state {
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
		if len(m.createmodel.Fields) == 0 {
			return "No fields to display."
		} else {

			current := m.createmodel.Fields[m.createmodel.EditLine]

			res += lipgloss.Place(
				width,
				height,
				lipgloss.Center,
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Center,
					current.Question,
					m.createmodel.View(),
				),
			)
		}
	}

	return res
}
