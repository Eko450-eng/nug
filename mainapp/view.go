package mainapp

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	s := ""
	if len(m.taskoverview.Tasks) <= 0 {
		s += "Coffee for all...."
	}

	switch m.state {
	case mainState:
		s += m.taskoverview.View(m.width, m.height)
	case calendarState:
		s += lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			m.calendar.View(m.width, m.height),
		)

	case helpState:
		s = m.helpmodal.View(m.width, m.height)
	}

	return s
}
