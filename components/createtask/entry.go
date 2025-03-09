package createtask

import (
	"nug/elements"
	"nug/structs"

	"github.com/charmbracelet/bubbles/textinput"
)

type CreateModel struct {
	EditLine          int
	Fields            []structs.Questions
	Newtask           structs.Task
	styles            structs.Styles
	Exiting           bool
	createInputFields textinput.Model
}

func InitTaskCreation() CreateModel {
	questions := []structs.Questions{
		elements.NewShortQuestion("Name"),
		elements.NewLongQuestion("Description"),
		elements.NewShortQuestion("Project_id"),
		elements.NewShortQuestion("Prio"),
		elements.NewShortQuestion("Completed"),
		elements.NewShortQuestion("Time"),
	}

	return CreateModel{
		EditLine: 0,
		Fields:   questions,
		Newtask:  structs.Task{},
		styles:   *structs.DefaultStyles(),
		Exiting:  false,
	}
}
