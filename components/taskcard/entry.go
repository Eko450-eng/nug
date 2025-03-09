package taskcard

import (
	"nug/elements"
	"nug/inputs"
	"nug/structs"
)

type Fields struct {
	Question   string
	Answer     string
	InputField inputs.Input
}

type TaskCardModel struct {
	Task     structs.Task
	Exiting  bool
	styles   structs.Styles
	Editing  bool
	editline int
	cursor   int
	fields   []structs.Questions
	current  structs.Questions
	IsActive bool
}

func InitModel(task structs.Task, isActive bool) TaskCardModel {
	fields := []structs.Questions{
		elements.NewShortQuestion("Name"),
		elements.NewLongQuestion("Description"),
		elements.NewShortQuestion("Project_id"),
		elements.NewShortQuestion("Prio"),
		elements.NewShortQuestion("Time"),
	}

	return TaskCardModel{
		Task:     task,
		cursor:   0,
		styles:   *structs.DefaultStyles(),
		fields:   fields,
		current:  fields[0],
		Exiting:  false,
		IsActive: isActive,
	}
}
