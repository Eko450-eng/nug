package taskoverview

import (
	"fmt"
	"nug/helpers"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) taskElementView() string {
	projects := helpers.GetProjects()

	borderStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#a7c957")).
		Foreground(lipgloss.Color("#000000"))

	taskText := "Prio - Task\n"

	for _, project := range projects {
		taskText += fmt.Sprintf("%s\n", borderStyle.Render(project.Name))

		for i, task := range m.Tasks {
			prio := lipgloss.NewStyle()
			normalTask := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff"))
			if task.Project_id == int(project.ID) {
				cursor := "  "
				if m.Cursor == i {
					cursor = "> "
				}

				if task.Completed == 1 && task.Deleted == 1 {
					normalTask = normalTask.
						Foreground(lipgloss.Color("9")).
						Strikethrough(true)
				} else if task.Completed == 1 {
					normalTask = normalTask.
						Strikethrough(true)
				} else if task.Deleted == 1 {
					normalTask = normalTask.
						Foreground(lipgloss.Color("9"))
				}

				switch task.Prio {
				case 2:
					prio = prio.Foreground(lipgloss.Color("#823038"))
				case 3:
					prio = prio.Foreground(lipgloss.Color("#96031A"))
				default:
					prio = prio.Foreground(lipgloss.Color("#B9E28C"))
				}

				taskText += fmt.Sprintf("%s%s %s\n",
					cursor,
					prio.Render(strconv.Itoa(task.Prio)),
					normalTask.Render(
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
