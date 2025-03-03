package main

import (
	"os"

	"nask/components/createtask"
	"nask/components/taskcard"
	"nask/components/taskoverview"
	"nask/helpers"
	"nask/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	mainState sessionState = iota
	createState
	taskState
)

type model struct {
	app_path     string
	tasks        []structs.Task
	show_deleted bool
	selected     map[int]struct{}
	width        int
	height       int
	styles       structs.Styles
	state        sessionState

	createmodel  createtask.CreateModel
	taskcard     taskcard.TaskCardModel
	taskoverview taskoverview.Model
}

func initModel() model {
	app_path, err := os.UserConfigDir()
	helpers.CheckErr(err)

	return model{
		app_path: app_path,
		selected: make(map[int]struct{}),
		tasks:    helpers.UpdateTasks(false),
		styles:   *structs.DefaultStyles(),

		state: mainState,

		createmodel:  createtask.CreateModel{},
		taskcard:     taskcard.TaskCardModel{},
		taskoverview: taskoverview.InitModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case createState:
		newState, newCmd := m.createmodel.UpdateCreateElement(msg)
		cmd = newCmd
		m.createmodel = newState
		if newState.Exiting {
			m.state = 0
			m.tasks = helpers.UpdateTasks(m.show_deleted)
		}
		return m, cmd
	case taskState:
		newState, newCmd := m.taskcard.UpdateTaskCard(msg)
		cmd = newCmd
		m.taskcard = newState
		if newState.ChangeEvent {
			m.tasks = helpers.UpdateTasks(m.show_deleted)
			m.taskcard.ChangeEvent = false
		}
		if newState.Exiting {
			m.state = 0
			m.tasks = helpers.UpdateTasks(m.show_deleted)
		}
		return m, cmd
	case mainState:
		newState, newCmd := m.taskoverview.Update(msg)
		m.taskoverview = newState
		cmd = newCmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
		}
		if key.Matches(msg, structs.Keymap.Create) {
			m.state = createState
			m.createmodel = createtask.InitTaskCreation()
		}
		if key.Matches(msg, structs.Keymap.Back) {
			m.state = mainState
		}
		if key.Matches(msg, structs.Keymap.Edit) {
			m.state = taskState
			m.taskcard.IsActive = true
			m.taskcard = taskcard.InitModel(m.tasks[m.taskoverview.Cursor], m.taskcard.IsActive)
		}
		return m, cmd
	}

	return m, cmd
}

func (m model) View() string {
	s := ""
	if len(m.tasks) <= 0 {
		s += "Coffee for all...."
	}

	width := m.width - 10

	leftWidth := width / 3
	rightWidth := width - leftWidth

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColor)
	borderStyleActive := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.BorderColorActive)

	switch m.state {
	case mainState:
		if len(m.tasks) > 0 {
			m.taskcard = taskcard.InitModel(m.tasks[m.taskoverview.Cursor], m.taskcard.IsActive)
		}

		s += lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyleActive.Width(leftWidth).Height(m.height-10).Render(
					m.taskoverview.View(),
				),
				borderStyle.Width(rightWidth).Height(m.height-10).Render(
					m.taskcard.View(rightWidth-30),
				),
			),
		)
	case taskState:
		s += lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyle.Width(leftWidth).Height(m.height-10).Render(
					m.taskoverview.View(),
				),
				borderStyleActive.Width(rightWidth).Height(m.height-10).Render(
					m.taskcard.View(rightWidth-30),
				),
			),
		)

	case createState:
		if len(m.createmodel.Fields) == 0 {
			return "No fields to display."
		}

		current := m.createmodel.Fields[m.createmodel.EditLine]

		s += lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				current.Question,
				m.createmodel.View(),
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
