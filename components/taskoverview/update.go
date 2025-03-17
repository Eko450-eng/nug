package taskoverview

import (
	"nug/components/createproject"
	"nug/components/createtask"
	"nug/components/projectview"
	"nug/components/quicknotes"
	"nug/components/settings"
	"nug/components/taskcard"
	"nug/helpers"
	"nug/structs"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) UpdateMain(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.initializing {
			m.Viewport = viewport.New(100, msg.Height)
			m.Viewport.YPosition = lipgloss.Height(m.taskElementView())
			m.Viewport.SetContent(m.taskElementView())
			m.initializing = false
		}
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.CreateProject) {
			m.createProject = createproject.InitModel()
			m.createProject.Form.Init()
			m.state = createProjectState
		} else if key.Matches(msg, structs.Keymap.HideCompleted) {
			m.hideCompleted = !m.hideCompleted
			m.Cursor = 0
			m.Tasks = m.UpdateTasks()
			m.Viewport.SetContent(m.taskElementView())
		} else if key.Matches(msg, structs.Keymap.Filter) {
			switch m.ordering {
			case prioAsc:
				m.ordering = prioDesc
				m.Tasks = m.UpdateTasks()
				m.Viewport.SetContent(m.taskElementView())
			case prioDesc:
				m.ordering = none
				m.Tasks = m.UpdateTasks()
				m.Viewport.SetContent(m.taskElementView())
			case none:
				m.ordering = prioAsc
				m.Tasks = m.UpdateTasks()
				m.Viewport.SetContent(m.taskElementView())
			}
		} else if key.Matches(msg, structs.Keymap.Up) && m.Cursor > 0 {
			m.Cursor--
			m.Viewport.LineUp(1)
			m.Viewport.SetContent(m.taskElementView())
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.Top) {
			m.Cursor = 0
			m.Viewport.GotoTop()
			m.Viewport.SetContent(m.taskElementView())
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.Bottom) {
			m.Cursor = len(m.Tasks) - 1
			m.Viewport.GotoBottom()
			m.Viewport.SetContent(m.taskElementView())
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.Down) && m.Cursor < len(m.Tasks)-1 {
			m.Cursor++
			m.Viewport.LineDown(1)
			m.Viewport.SetContent(m.taskElementView())
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.ShowDeleted) {
			m.Cursor = 0
			m.show_deleted = !m.show_deleted
			m.Tasks = m.UpdateTasks()
			m.Viewport.SetContent(m.taskElementView())
		} else if key.Matches(msg, structs.Keymap.Delete) {
			db, _ := helpers.ConnectToSQLite()
			currenttask := m.Tasks[m.Cursor]
			if currenttask.Deleted == 1 {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.ID).
					Update("Deleted", 0)
				m.Tasks = m.UpdateTasks()
				m.Viewport.SetContent(m.taskElementView())
				m.Cursor = 0
			} else {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.ID).
					Update("Deleted", 1)
				m.Tasks = m.UpdateTasks()
				m.Viewport.SetContent(m.taskElementView())
				if m.Cursor == 0 {
					m.Cursor = 0
				} else {
					m.Cursor--
				}
				if len(m.Tasks) > 0 {
					m.taskcard.Task = m.Tasks[m.Cursor]
				}
			}
		} else if key.Matches(msg, structs.Keymap.Check) {
			db, _ := helpers.ConnectToSQLite()
			currenttask := m.Tasks[m.Cursor]
			newvalue := 0
			if currenttask.Completed == 1 {
				newvalue = 0
			} else {
				newvalue = 1
			}
			db.Model(&structs.Task{}).
				Where("id = ?", currenttask.ID).
				Update("Completed", newvalue)
			m.Tasks = m.UpdateTasks()
			m.Viewport.SetContent(m.taskElementView())
		} else if key.Matches(msg, structs.Keymap.Edit) {
			m.taskcard = taskcard.InitModel(m.Tasks[m.Cursor], false)
			m.taskcard.Form.Init()
			m.taskcard.IsActive = true
			m.taskcard.Task = m.Tasks[m.Cursor]

			m.state = infoState
		} else if key.Matches(msg, structs.Keymap.Create) {
			m.state = createState
			m.createmodel = createtask.InitTaskCreation()
			m.createmodel.Form.Init()
		}
	}
	return m, cmd
}

func (m Model) UpdateTask(msg tea.Msg) (taskcard.TaskCardModel, tea.Cmd) {
	var cmd tea.Cmd

	newState, newCmd := m.taskcard.UpdateTaskCard(msg)
	cmd = newCmd
	m.taskcard = newState

	return newState, cmd
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case quickNoteState:
		newState, newCmd := m.quickNote.Update(msg)
		m.quickNote = newState
		cmd = newCmd

		if newState.Finished {
			m.state = mainState
			m.Tasks = m.UpdateTasks()
		}
		return m, cmd
	case projectViewState:
		newState, newCmd := m.projectView.Update(msg)
		m.projectView = newState
		cmd = newCmd

		if newState.Finished {
			m.state = mainState
			m.Tasks = m.UpdateTasks()
		}
		return m, cmd
	case calendarState:
		newState, newCmd := m.calendar.Update(msg)
		m.calendar = newState
		cmd = newCmd
	case createState:
		newState, newCmd := m.createmodel.UpdateCreateElement(msg)
		cmd = newCmd
		m.createmodel = newState
		if newState.Exiting {
			m.state = 0
			m.Tasks = m.UpdateTasks()
		}
		return m, cmd

	case mainState:
		newState, newCmd := m.UpdateMain(msg)
		m = newState
		cmd = newCmd

	case settingState:
		newState, newCmd := m.settings.Update(msg)
		m.settings = newState
		cmd = newCmd

		if newState.Finished {
			m.state = mainState
			m.Tasks = m.UpdateTasks()
		}
		return m, cmd
	case createProjectState:
		newState, newCmd := m.createProject.Update(msg)
		m.createProject = newState
		cmd = newCmd

		if newState.Finished {
			m.state = mainState
			m.Tasks = m.UpdateTasks()
		}
		return m, cmd
	case infoState:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if key.Matches(msg, structs.Keymap.Quit) {
				cmd = tea.Quit
			}
		}

		newState, newCmd := m.UpdateTask(msg)
		m.taskcard = newState

		if m.taskcard.Exiting {
			m.Tasks = m.UpdateTasks()
			m.state = mainState

			m.taskcard.Task = m.Tasks[m.Cursor]

			m.taskcard.IsActive = false
			m.taskcard.Exiting = false
		}

		cmd = newCmd
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.QuickNotes) {
			switch m.state {
			case mainState:
				m.quickNote = quicknotes.InitModel()
				m.quickNote.Form.Init()
				m.state = quickNoteState
			}
		} else if key.Matches(msg, structs.Keymap.Settings) {
			switch m.state {
			case projectViewState:
				m.state = mainState
			case settingState:
				m.state = mainState
			case mainState:
				m.settings = settings.InitModel()
				m.settings.Form.Init()
				m.state = settingState
			}
		} else if key.Matches(msg, structs.Keymap.TabSwitch) {
			switch m.state {
			case mainState:
				m.calendar.Selected = time.Now().Day() - 1
				m.calendar.HideCompleted = m.hideCompleted
				m.state = calendarState
			case calendarState:
				m.projectView = projectview.InitModel()
				m.projectView.Form.Init()
				m.state = projectViewState
			}
		} else if key.Matches(msg, structs.Keymap.Sync) {
			helpers.SyncToWebDav()
		} else if key.Matches(msg, key.NewBinding(key.WithKeys("r"))) {
			m.Tasks = m.UpdateTasks()
		}
	}

	return m, cmd
}
