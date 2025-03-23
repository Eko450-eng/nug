package quicknoteview

import (
	"nug/structs"

	"github.com/charmbracelet/huh"
)

type Model struct {
	Finished bool
	Note     []structs.QuickNotes
	Form     *huh.Form
}

func InitModel() Model {
	return Model{
		Finished: false,
		Note:     []structs.QuickNotes{},
		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("New Quicke Note").
					Key("note"),
			),
		),
	}
}
