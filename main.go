package main

import (
	"fmt"
	"os"

	"nask/components"
	"nask/components/taskview"
	"nask/helpers"
	"nask/structs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	app_path     string
	tasks        []structs.Task
	show_deleted bool
	cursor       int
	selected     map[int]struct{}
	width        int
	height       int
	styles       structs.Styles
	createmodel  components.CreateModel
	taskview     taskview.TaskViewModel
	creating     bool
	editing      bool
}

func initModel() model {
	app_path, err := os.UserConfigDir()
	helpers.CheckErr(err)

	return model{
		app_path:    app_path,
		cursor:      0,
		selected:    make(map[int]struct{}),
		tasks:       helpers.UpdateTasks(false),
		creating:    false,
		editing:     false,
		createmodel: components.CreateModel{},
		taskview:    taskview.TaskViewModel{},
		styles:      *structs.DefaultStyles(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.creating {
	case true:
		newState, newCmd := m.createmodel.UpdateCreateElement(msg)
		cmd = newCmd
		m.createmodel = newState
		if newState.Exiting {
			m.creating = false
			m.tasks = helpers.UpdateTasks(m.show_deleted)
		}
		return m, cmd
	case false:
		switch m.editing {
		case true:
			newState, newCmd := m.taskview.Update(msg)
			cmd = newCmd
			m.taskview = newState
			if newState.ChangeEvent {
				m.tasks = helpers.UpdateTasks(m.show_deleted)
				m.taskview.ChangeEvent = false
			}
			if newState.Exiting {
				m.editing = false
				m.tasks = helpers.UpdateTasks(m.show_deleted)
			}
			return m, cmd
		case false:
			switch msg := msg.(type) {
			case tea.WindowSizeMsg:
				m.width = msg.Width
				m.height = msg.Height
			case tea.KeyMsg:
				switch msg.String() {
				case "ctrl+c":
					cmd = tea.Quit
				case "c":
					m.creating = true
					m.createmodel = components.InitTaskCreation()
				case "esc":
					m.creating = false
				case "r":
					m.tasks = helpers.UpdateTasks(m.show_deleted)
				case "k", "up":
					if m.cursor > 0 {
						m.cursor--
					}
				case "j", "down":
					if m.cursor < len(m.tasks)-1 {
						m.cursor++
					}
				case "d":
					m.cursor = 0
					m.show_deleted = !m.show_deleted
					m.tasks = helpers.UpdateTasks(m.show_deleted)
				case "D":
					db, _ := helpers.ConnectToSQLite()
					currenttask := m.tasks[m.cursor]
					if currenttask.Deleted == 1 {
						db.Model(&structs.Task{}).
							Where("id = ?", currenttask.Id).
							Update("Deleted", 0)
						m.tasks = helpers.UpdateTasks(false)
						m.show_deleted = false
						m.cursor = 0
					} else {
						db.Model(&structs.Task{}).
							Where("id = ?", currenttask.Id).
							Update("Deleted", 1)
						m.tasks = helpers.UpdateTasks(m.show_deleted)
						if m.cursor == 0 {
							m.cursor = 0
						} else {
							m.cursor--
						}
					}

				case " ":
					db, _ := helpers.ConnectToSQLite()
					currenttask := m.tasks[m.cursor]
					newvalue := 0
					if currenttask.Completed == 1 {
						newvalue = 0
					} else {
						newvalue = 1
					}
					db.Model(&structs.Task{}).
						Where("id = ?", currenttask.Id).
						Update("Completed", newvalue)
					m.tasks = helpers.UpdateTasks(m.show_deleted)
				case "enter", "l":
					m.editing = true
					m.taskview = taskview.InitTaskCreation(m.tasks[m.cursor], true)
				}
			}

		}
	}

	return m, cmd
}

func (m model) View() string {
	s := ""
	if len(m.tasks) <= 0 {
		s += "Coffee for all...."
	}

	tasks := ""

	leftWidth := m.width / 3
	rightWidth := m.width - leftWidth

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(m.styles.BorderColor)

	if !m.creating {
		for i, task := range m.tasks {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if m.tasks[i].Completed == 1 {
				checked = "x"
			}
			if m.tasks[i].Deleted == 1 {
				tasks += fmt.Sprintf("%s D-[%s] %s\n", cursor, checked, task.Name)
			} else {
				tasks += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Name)
			}
		}

		if !m.editing && len(m.tasks) > 0 {
			m.taskview = taskview.InitTaskCreation(m.tasks[m.cursor], false)
		}

		s = lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyle.Width(leftWidth).Height(m.height-10).Render(
					tasks,
				),
				borderStyle.Width(rightWidth).Height(m.height-10).Render(
					m.taskview.View(rightWidth-30),
				),
			),
		)
	} else {
		createModelView := m.createmodel.View()

		if len(m.createmodel.Fields) == 0 {
			return "No fields to display."
		}

		current := m.createmodel.Fields[m.createmodel.EditLine]

		s = lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				current.Question,
				createModelView,
			),
		)
	}

	return s
}

func main() {
	db, err := helpers.ConnectToSQLite()
	db.AutoMigrate(&structs.Task{})
	helpers.CheckErr(err)

	// Start the app
	f, err := tea.LogToFile("debug.log", "debug")
	helpers.CheckErr(err)

	defer f.Close()
	p := tea.NewProgram(initModel())
	_, err = p.Run()
	helpers.CheckErr(err)
}
