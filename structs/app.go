package structs

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	Up            key.Binding
	Down          key.Binding
	Left          key.Binding
	Right         key.Binding
	Save          key.Binding
	Create        key.Binding
	CreateProject key.Binding
	Edit          key.Binding
	Check         key.Binding
	ShowDeleted   key.Binding
	Delete        key.Binding
	SkipForm      key.Binding
	Back          key.Binding
	Quit          key.Binding
	QuitEasy      key.Binding
	Help          key.Binding
	TabSwitch     key.Binding
	Filter        key.Binding
	Sync          key.Binding
	HideCompleted key.Binding
	Settings      key.Binding
	QuickNotes    key.Binding
	Top           key.Binding
	Bottom        key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	QuickNotes: key.NewBinding(key.WithKeys("E"), key.WithHelp("Shift + E", "Add Quicknote")),
	Top:        key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "Top")),
	Bottom:     key.NewBinding(key.WithKeys("G"), key.WithHelp("Shift + G", "Bottom")),
	Settings: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "Show settings"),
	),
	HideCompleted: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "Hide or show Completed"),
	),
	Sync: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "Sync to WebDav"),
	),
	Filter: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "Filter"),
	),
	TabSwitch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("Tab", "Switch to Calendar and back"),
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
	CreateProject: key.NewBinding(
		key.WithKeys("ctrl+e"),
		key.WithHelp("ctrl+e", "New Project"),
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
		key.WithKeys("ctrl+q"),
		key.WithHelp("ctrl+q", "quit"),
	),
	QuitEasy: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quitEasy"),
	),
}
