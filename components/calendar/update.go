package calendar

import (
	"nug/structs"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	month := time.Now().Local().Month()
	_, r := 2025, month

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Up) && m.Selected < DaysInMonth(time.Now().Year(), r) {
			m.Selected -= 7
		} else if key.Matches(msg, structs.Keymap.Down) && m.Selected > 0 {
			m.Selected += 7
		} else if key.Matches(msg, structs.Keymap.Right) && m.Selected < DaysInMonth(time.Now().Year(), r) {
			m.Selected++
		} else if key.Matches(msg, structs.Keymap.Left) && m.Selected > 0 {
			m.Selected--
		}
	}

	return m, cmd
}
