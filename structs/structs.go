package structs

import (
	"nug/inputs"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

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

type Task struct {
	gorm.Model
	Id          int `gorm:"primaryKey`
	Name        string
	Description string
	Project_id  int
	Prio        int
	Time        string
	Date        string
	Deletedtime string
	Modified    string
	Completed   int
	Deleted     int
}

type Tags struct {
	gorm.Model
	Id          int `gorm:"primaryKey`
	Name        string
	Deleted     int
	Deletedtime string
}

type Tag_to_task struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Tag  int
	Task int
}

type Project struct {
	gorm.Model
	Id          int `gorm:"primaryKey`
	Name        string
	Deleted     int
	Deletedtime string
}

type Task_to_Project struct {
	gorm.Model
	Id      int `gorm:"primaryKey"`
	Task    int
	Project int
}

type Questions struct {
	Question   string
	Answer     string
	InputField inputs.Input
}

type keymap struct {
	Up          key.Binding
	Down        key.Binding
	Left        key.Binding
	Right       key.Binding
	Save        key.Binding
	Create      key.Binding
	Edit        key.Binding
	Check       key.Binding
	ShowDeleted key.Binding
	Delete      key.Binding
	SkipForm    key.Binding
	Back        key.Binding
	Quit        key.Binding
	Help        key.Binding
	TabSwitch   key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	TabSwitch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("Tab", "Switch focus"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Help"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("left/h", "Left"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("down/j", "down"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("up/k", "Up"),
	),
	Edit: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "Edit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "h"),
		key.WithHelp("esc/h", "back"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("right/l", "Right"),
	),
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "New"),
	),
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+enter", "Save"),
	),
	Check: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "Check"),
	),
	Delete: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "Delete/Restore"),
	),
	ShowDeleted: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "Toggle Deleted"),
	),
	SkipForm: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "QuickSave"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
