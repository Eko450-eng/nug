package structs

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor       lipgloss.Color
	BorderColorActive lipgloss.Color
	InputField        lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("#2E5077")
	s.BorderColorActive = lipgloss.Color("#F77F00")
	s.InputField = lipgloss.
		NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(80)
	return s
}
