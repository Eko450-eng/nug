package taskoverview

import (
	"fmt"
	"nug/helpers"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) taskElementView() string {
	projects := helpers.GetProjects()
	checked := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff"))

	borderStyle := checked.
		Background(lipgloss.Color("#a7c957")).
		Foreground(lipgloss.Color("#000000"))

	prio := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9"))

	taskText := "Prio - Task\n"

	cursor := "  "

	for _, project := range projects {
		taskText += fmt.Sprintf("%s\n", borderStyle.Render(project.Name))

		for i, task := range m.Tasks {
			if task.Project_id == int(project.ID) {
				if task.Completed == 1 {
					checked = checked.
						Foreground(lipgloss.Color("#9EADC8")).
						Strikethrough(true)
				}

				if m.Cursor == i {
					checked = checked.Bold(true)
					cursor = "> "
				}

				if task.Deleted == 1 {
					checked = checked.
						Foreground(lipgloss.Color("9"))
					if task.Completed == 1 {
						checked.Strikethrough(true)
					}
				}

				taskText += fmt.Sprintf("%s\n",
					checked.Render(cursor,
						prio.Render(strconv.Itoa(task.Prio)),
						helpers.ShortenString(task.Name, 30),
					))
			}
		}
	}

	return taskText
}

func (m Model) View(width, height int) string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColor)

	borderStyleActive := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColorActive)

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
				m.Viewport.View(),
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
					m.Viewport.View(),
				),
				borderStyleActive.Width(rightWidth).Height(height).Render(
					m.taskcard.View(rightWidth-30),
				),
			),
		)
	case createState:
		m.createmodel.Init()

		res += lipgloss.Place(
			width,
			height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.createmodel.View(),
			),
		)
	}

	return res
}
