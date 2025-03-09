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
	current           structs.Questions
}

func InitTaskCreation() CreateModel {
	questions := []structs.Questions{
		elements.NewShortQuestion("Name"),
		elements.NewLongQuestion("Description"),
		elements.NewShortQuestion("Project_id"),
		elements.NewShortQuestion("Prio"),
	}

	return CreateModel{
		EditLine: 0,
		Fields:   questions,
		current:  questions[0],
		Newtask:  structs.Task{},
		styles:   *structs.DefaultStyles(),
		Exiting:  false,
	}
}
