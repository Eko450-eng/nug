package helpers

import (
	"fmt"
	"os"

	"nug/structs"

	"github.com/charmbracelet/huh"
	"github.com/glebarez/sqlite"
	gap "github.com/muesli/go-app-paths"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	scope := gap.NewScope(gap.User, "nug")
	dirs, err := scope.DataDirs()
	appPath := dirs[0] + "/nug.db"
	CheckErr(err)
	dirPath := dirs[0]

	_, statErr := os.Stat(dirPath)
	if os.IsNotExist(statErr) {
		err := os.MkdirAll(dirPath, 0o770)
		CheckErr(err)
	}

	_, err = os.Stat(appPath)
	if os.IsNotExist(err) {
		_, err := os.Create(appPath)
		if err != nil {
			LogToFile(fmt.Sprintf("%v", err))
			return nil, fmt.Errorf("failed to create the database file: %v", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(appPath), &gorm.Config{})
	CheckErr(err)

	var tasks []structs.Task
	db.Find(&tasks)
	if len(tasks) <= 0 {
		NewTask := structs.Task{
			Name:        "Welcome",
			Description: "You can delete this task now if you'd like :)",
			Project_id:  0,
			Prio:        0,
			Time:        "",
			Deletedtime: "",
			Modified:    "",
			Completed:   0,
			Deleted:     0,
		}
		db.Create(&NewTask)
	}

	return db, nil
}

func Resettask() structs.Task {
	return structs.Task{
		Name:        "",
		Description: "",
		Project_id:  0,
		Prio:        0,
		Time:        "",
		Deletedtime: "",
		Modified:    "",
		Completed:   0,
		Deleted:     0,
	}
}

func GetNotes() []structs.QuickNotes {
	var notes []structs.QuickNotes
	db, _ := ConnectToSQLite()
	if res := db.
		Find(&notes); res.Error != nil {
		panic(res.Error)
	}
	return notes
}

func GetProjectName(id int) string {
	var projects structs.Project
	db, _ := ConnectToSQLite()
	if res := db.
		Where("id = ?", id).
		Find(&projects); res.Error != nil {
		panic(res.Error)
	}
	return projects.Name
}

func GetSettings() structs.Settings {
	var settings structs.Settings
	db, _ := ConnectToSQLite()
	if res := db.
		Where("id = 1").
		Find(&settings); res.Error != nil {
		panic(res.Error)
	}
	return settings
}

func GetProjects() []structs.Project {
	var projects []structs.Project
	db, _ := ConnectToSQLite()
	if res := db.
		Find(&projects); res.Error != nil {
		panic(res.Error)
	}
	return projects
}

func GetTaskCountOfProject(projectId uint) int {
	var tasks []structs.Task
	db, _ := ConnectToSQLite()

	if res := db.
		Where("project_id = ? AND deleted = 0 AND completed = 0", projectId).
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}

	return len(tasks)
}

func GetProjectItems() []huh.Option[int] {
	projects := GetProjects()
	projectItems := []huh.Option[int]{}

	for _, project := range projects {
		projectItems = append(projectItems, huh.NewOption(project.Name, int(project.ID)))
	}

	return projectItems
}

func GetFilteredTask(prio, projectid int) []structs.Task {
	var tasks []structs.Task
	db, _ := ConnectToSQLite()
	if res := db.
		Where("deleted = ?", 0).
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}
	return tasks
}
