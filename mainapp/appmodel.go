package mainapp

import (
	"os"

	"nug/components/calendar"
	"nug/components/helpmodal"
	"nug/components/taskoverview"
	"nug/helpers"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	mainState sessionState = iota
	createState
	helpState
	calendarState
)

type model struct {
	app_path string
	width    int
	height   int
	styles   structs.Styles
	state    sessionState

	taskoverview taskoverview.Model
	helpmodal    helpmodal.Model
	calendar     calendar.Model
}

func InitModel() model {
	app_path, err := os.UserConfigDir()
	helpers.CheckErr(err)

	return model{
		app_path: app_path,
		styles:   *structs.DefaultStyles(),

		state: mainState,

		taskoverview: taskoverview.InitModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
