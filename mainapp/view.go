package mainapp

import "github.com/charmbracelet/lipgloss"

var mainStyle = lipgloss.NewStyle()

func (m model) View() string {
	s := ""

	switch m.state {
	case mainState:
		s = mainStyle.Render(m.taskoverview.View(m.width, m.height))

	case helpState:
		s = mainStyle.Render(m.helpmodal.View(m.width, m.height))
	}

	return s
}
