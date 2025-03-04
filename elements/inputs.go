package elements

import (
	"nug/inputs"
	"nug/structs"
)

func NewQuestion(q string) structs.Questions {
	return structs.Questions{Question: q}
}

func NewShortQuestion(q string) structs.Questions {
	question := NewQuestion(q)
	model := inputs.NewShortAnswerField()
	question.InputField = model
	return question
}

func NewLongQuestion(q string) structs.Questions {
	question := NewQuestion(q)
	model := inputs.NewLongAnswerField()
	question.InputField = model
	return question
}
