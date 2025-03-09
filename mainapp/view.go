package mainapp

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	s := ""
	if len(m.taskoverview.Tasks) <= 0 {
		s += "Coffee for all...."
	}

	width := m.width - 10

	leftWidth := width / 3
	rightWidth := width - leftWidth

	switch m.state {
	case mainState:
		s += m.taskoverview.View(m.width, m.height)
	case calendarState:
		s += lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			m.calendar.View(rightWidth-30, m.height),
		)

	case helpState:
		s = m.helpmodal.View(m.width, m.height)
	}

	return s
}
