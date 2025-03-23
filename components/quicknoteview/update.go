package quicknoteview

import (
	"nug/helpers"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, structs.Keymap.QuitEasy) {
			m.Finished = true
			return m, nil
		}
	}

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	if m.Form.State == huh.StateCompleted {
		db, _ := helpers.ConnectToSQLite()

		// note := m.Form.GetString("note")

		// note := structs.QuickNotes{
		// 	Note:    note,
		// 	Deleted: 0,
		// }

		db.Create(&m.Note)

		m.Finished = true
	}

	return m, cmd
}
