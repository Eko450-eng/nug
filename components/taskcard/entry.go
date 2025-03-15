package taskcard

import (
	"nug/elements"
	"nug/helpers"
	"nug/inputs"
	"nug/structs"

	"github.com/charmbracelet/huh"
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

	Form *huh.Form
}

func InitModel(task structs.Task, isActive bool) TaskCardModel {
	fields := []structs.Questions{
		elements.NewShortQuestion("Name"),
		elements.NewLongQuestion("Description"),
		elements.NewShortQuestion("Project_id"),
		elements.NewShortQuestion("Prio"),
		elements.NewShortQuestion("Date"),
	}
	projectItems := helpers.GetProjectItems()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("Name").
				Value(&task.Name),
			huh.NewText().
				Key("description").
				Title("Description").
				Value(&task.Description),
			huh.NewSelect[int]().
				Key("prio").
				Title("Priority").
				Value(&task.Prio).
				Options(
					huh.NewOption("1", 1),
					huh.NewOption("2", 2),
					huh.NewOption("3", 3),
				),
			huh.NewSelect[int]().
				Key("project").
				Title("Project").
				Value(&task.Project_id).
				Options(
					projectItems...,
				),
			huh.NewText().
				Key("date").
				Title("Date").
				Value(&task.Date),
		),
	)

	return TaskCardModel{
		Task:     task,
		cursor:   0,
		styles:   *structs.DefaultStyles(),
		fields:   fields,
		Exiting:  false,
		IsActive: isActive,
		Form:     form,
	}
}
