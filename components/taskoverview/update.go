package taskoverview

import (
	"nug/components/createtask"
	"nug/components/taskcard"
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) UpdateMain(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Up) && m.Cursor > 0 {
			m.Cursor--
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.Down) && m.Cursor < len(m.Tasks)-1 {
			m.Cursor++
			m.taskcard.Task = m.Tasks[m.Cursor]
		} else if key.Matches(msg, structs.Keymap.ShowDeleted) {
			m.Cursor = 0
			m.show_deleted = !m.show_deleted
			m.Tasks = m.UpdateTasks()
		} else if key.Matches(msg, structs.Keymap.Delete) {
			db, _ := helpers.ConnectToSQLite()
			currenttask := m.Tasks[m.Cursor]
			if currenttask.Deleted == 1 {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.Id).
					Update("Deleted", 0)
				m.Tasks = m.UpdateTasks()
				m.Cursor = 0
			} else {
				db.Model(&structs.Task{}).
					Where("id = ?", currenttask.Id).
					Update("Deleted", 1)
				m.Tasks = m.UpdateTasks()
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
				Where("id = ?", currenttask.Id).
				Update("Completed", newvalue)
			m.Tasks = m.UpdateTasks()
		} else if key.Matches(msg, structs.Keymap.Edit) {
			m.taskcard.IsActive = true
			m.taskcard.Task = m.Tasks[m.Cursor]

			m.state = infoState
		} else if key.Matches(msg, structs.Keymap.Create) {
			m.state = createState
			m.createmodel = createtask.InitTaskCreation()
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Sync) {
			helpers.SyncToWebDav()
		} else if key.Matches(msg, key.NewBinding(key.WithKeys("r"))) {
			m.Tasks = m.UpdateTasks()
		}
	}

	switch m.state {
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

	case infoState:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if key.Matches(msg, structs.Keymap.Quit) {
				cmd = tea.Quit
			} else if key.Matches(msg, structs.Keymap.Back) && !m.taskcard.Editing {
				m.Tasks = m.UpdateTasks()
				m.state = mainState

				m.taskcard.Task = m.Tasks[m.Cursor]

				m.taskcard.IsActive = false
				m.taskcard.Exiting = false
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
	}

	return m, cmd
}
