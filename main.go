package main

import (
	"nug/helpers"
	"nug/mainapp"
	"nug/structs"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	db, err := helpers.ConnectToSQLite()
	db.AutoMigrate(&structs.Task{})
	db.AutoMigrate(&structs.Project{})
	db.AutoMigrate(&structs.Settings{})
	helpers.CheckErr(err)

	f, err := tea.LogToFile("debug.log", "debug")
	helpers.CheckErr(err)

	defer f.Close()
	p := tea.NewProgram(mainapp.InitModel())
	_, err = p.Run()
	helpers.CheckErr(err)
}
