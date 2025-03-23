package quicknoteview

import (
	"nug/helpers"

	"github.com/charmbracelet/lipgloss"
)

var mainStyle = lipgloss.NewStyle()
var borderStyle = mainStyle.
	Background(lipgloss.Color("#a7c957")).
	Foreground(lipgloss.Color("#000000"))

func (m Model) View() string {
	s := "Notes\n"

	for _, note := range m.Note {
		timeStamp := note.CreatedAt
		t := timeStamp.Format("02.01.2006 15:04")

		helpers.LogToFile(t)

		s += mainStyle.Render(
			t,
			"\n",
			borderStyle.Render(
				note.Note,
			),
			"\n",
			"\n",
		)
	}

	return s
}
