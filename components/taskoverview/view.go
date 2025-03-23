package taskoverview

import (
	"fmt"
	"nug/helpers"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var mainStyle = lipgloss.NewStyle()

func (m Model) taskElementView() string {
	projects := helpers.GetProjects()

	borderStyle := mainStyle.
		Background(lipgloss.Color("#a7c957")).
		Foreground(lipgloss.Color("#000000"))

	taskText := "Prio - Task\n"

	for _, project := range projects {
		taskText += fmt.Sprintf("%s\n", borderStyle.Render(project.Name))

		for i, task := range m.Tasks {
			prio := mainStyle
			normalTask := mainStyle.
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

	return mainStyle.Render(taskText)
}

func (m Model) View(width, height int) string {
	borderStyle := mainStyle.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColor)

	borderStyleActive := mainStyle.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColorActive)

	width = width - 10
	height = height - 2

	leftWidth := width / 3
	rightWidth := width - leftWidth

	res := ""

	switch m.state {
	case quickNoteViewState:
		res = lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			mainStyle.Render(
				m.quicknoteview.View(),
			),
		)
	case settingState:
		res = lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			mainStyle.Render(
				m.settings.View(),
			),
		)
	case calendarState:
		res = lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			mainStyle.Render(
				m.calendar.View(width, height),
			),
		)
	case mainState:
		res = lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			mainStyle.Render(
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					m.Viewport.View(),
					borderStyle.Width(rightWidth).Height(height).Render(
						m.taskcard.View(rightWidth-30),
					),
				),
			),
		)

	case projectViewState:
		res = mainStyle.Render(m.projectView.View())
	case createProjectState:
		res = mainStyle.Render(m.createProject.View())
	case quickNoteState:
		res = mainStyle.Render(m.quickNote.View())

	case infoState:
		res = lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			mainStyle.Render(
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					borderStyle.Width(leftWidth).Height(height).Render(
						m.Viewport.View(),
					),
					borderStyleActive.Width(rightWidth).Height(height).Render(
						m.taskcard.View(rightWidth-30),
					),
				),
			),
		)
	case createState:
		m.createmodel.Init()

		res = lipgloss.Place(
			width,
			height,
			lipgloss.Center,
			lipgloss.Center,
			mainStyle.Render(
				lipgloss.JoinVertical(
					lipgloss.Center,
					m.createmodel.View(),
				),
			),
		)
	}

	return mainStyle.Render(res)
}
