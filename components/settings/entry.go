package settings

import (
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type orderState int

const (
	prioAsc orderState = iota
	prioDesc
	none
)

type Model struct {
	Finished bool
	Form     *huh.Form

	settings structs.Settings
}

func InitModel() Model {
	settings := helpers.GetSettings()

	return Model{
		Finished: false,

		settings: structs.Settings{
			HideCompleted: settings.HideCompleted,
			Ordering:      settings.Ordering,
		},

		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Key("hidecompleted").
					Title("Hide Completed").
					Value(&settings.HideCompleted).
					Options(
						huh.NewOption("Don't hide", 0),
						huh.NewOption("Hide", 1),
					),
				huh.NewSelect[int]().
					Key("ordering").
					Title("Ordering").
					Value(&settings.Ordering).
					Options(
						huh.NewOption("Prio ASC", 0),
						huh.NewOption("Prio Desc", 1),
						huh.NewOption("Default (By ID)", 2),
					),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Form.Init()
}
