package helpmodal

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var mainStyle = lipgloss.NewStyle()

func (m Model) View(width, height int) string {
	borderStyle := mainStyle.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			borderStyle.Render(
				m.Table.View(),
			),
		),
	)

}
