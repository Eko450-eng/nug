package taskcard

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func selectedCursor(cursor, pos int, isActive bool) string {
	c := ""
	if cursor == pos && isActive {
		c = ">"
	}
	return c
}

func (m TaskCardModel) View(width int) string {
	borderColor := m.styles.BorderColor
	if m.Task.Deleted == 1 {
		borderColor = lipgloss.Color("9")
	}
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(width)

	if m.Editing {
		return m.styles.InputField.Render(m.current.InputField.View())
	} else {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			borderStyle.Render(
				fmt.Sprintf("Last Updated: %s", m.Task.UpdatedAt.Format("02.01.2006 at 15:04")),
			),

			borderStyle.Render(
				fmt.Sprintf("Name:\n%s %s", selectedCursor(m.cursor, 0, m.IsActive), m.Task.Name),
			),
			borderStyle.Render(
				fmt.Sprintf("Description:\n%s %s", selectedCursor(m.cursor, 1, m.IsActive), m.Task.Description),
			),
			borderStyle.Render(
				fmt.Sprintf("Project_id:\n%s %s", selectedCursor(m.cursor, 2, m.IsActive), strconv.Itoa(m.Task.Project_id)),
			),
			borderStyle.Render(
				fmt.Sprintf("Prio:\n%s %s", selectedCursor(m.cursor, 3, m.IsActive), strconv.Itoa(m.Task.Prio)),
			),
			borderStyle.Render(
				fmt.Sprintf("Time:\n%s %s", selectedCursor(m.cursor, 4, m.IsActive), m.Task.Time),
			),
		)
	}
}
