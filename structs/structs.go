package structs

import (
	"nask/inputs"

	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.
		NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(80)
	return s
}

type Task struct {
	gorm.Model
	Id          int `gorm:"primaryKey`
	Name        string
	Description string
	Project_id  int
	Prio        int
	Time        string
	Deletedtime string
	Modified    string
	Completed   int
	Deleted     int
}

type Questions struct {
	Question   string
	Answer     string
	InputField inputs.Input
}
