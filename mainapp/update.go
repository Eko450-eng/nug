package mainapp

import (
	"nug/components/helpmodal"
	"nug/structs"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case helpState:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if key.Matches(msg, structs.Keymap.Quit) {
				cmd = tea.Quit
				return m, cmd
			} else {
				m.state = mainState
				return m, cmd
			}
		}
	case mainState:
		newState, newCmd := m.taskoverview.Update(msg)
		m.taskoverview = newState
		cmd = newCmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.taskoverview.Width = m.width
		m.taskoverview.Height = m.height
	case tea.KeyMsg:
		if key.Matches(msg, structs.Keymap.Help) {
			m.helpmodal = helpmodal.Init()
			m.state = helpState
		} else if key.Matches(msg, structs.Keymap.Quit) {
			cmd = tea.Quit
		}
		return m, cmd
	}

	return m, cmd
}
