package helpmodal

import (
	"nug/structs"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Table table.Model
}

func Init() Model {
	cols := []table.Column{
		{Title: "Keybind", Width: 10},
		{Title: "Description", Width: 20},
	}

	rows := []table.Row{
		{structs.Keymap.Up.Help().Key, structs.Keymap.Up.Help().Desc},
		{structs.Keymap.Down.Help().Key, structs.Keymap.Down.Help().Desc},
		{structs.Keymap.Left.Help().Key, structs.Keymap.Left.Help().Desc},
		{structs.Keymap.Right.Help().Key, structs.Keymap.Right.Help().Desc},
		{structs.Keymap.Save.Help().Key, structs.Keymap.Save.Help().Desc},
		{structs.Keymap.Create.Help().Key, structs.Keymap.Create.Help().Desc},
		{structs.Keymap.Edit.Help().Key, structs.Keymap.Edit.Help().Desc},
		{structs.Keymap.Check.Help().Key, structs.Keymap.Check.Help().Desc},
		{structs.Keymap.ShowDeleted.Help().Key, structs.Keymap.ShowDeleted.Help().Desc},
		{structs.Keymap.Delete.Help().Key, structs.Keymap.Delete.Help().Desc},
		{structs.Keymap.SkipForm.Help().Key, structs.Keymap.SkipForm.Help().Desc},
		{structs.Keymap.Back.Help().Key, structs.Keymap.Back.Help().Desc},
		{structs.Keymap.Quit.Help().Key, structs.Keymap.Quit.Help().Desc},
		{structs.Keymap.TabSwitch.Help().Key, structs.Keymap.TabSwitch.Help().Desc},
		{structs.Keymap.Sync.Help().Key, structs.Keymap.Sync.Help().Desc},
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	return Model{
		Table: t,
	}
}
