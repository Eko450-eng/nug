package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"nask/structs"

	"github.com/glebarez/sqlite"
	gap "github.com/muesli/go-app-paths"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	scope := gap.NewScope(gap.User, "nask")
	dirs, err := scope.DataDirs()
	appPath := dirs[0] + "/nask.db"
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

	// Create DB
	db.AutoMigrate(&structs.Task{})

	var tasks []structs.Task
	db.Find(&tasks)
	if len(tasks) <= 0 {
		NewTask := structs.Task{
			Id:          0,
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

func SetDefaultInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		num = 0
	}
	return num
}

func CheckErr(err error) {
	if err != nil {
		LogToFile(fmt.Sprintf("%s", err))
	}
}

func Resettask() structs.Task {
	return structs.Task{
		Id:          0,
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

func UpdateTasks(show_deleted bool) []structs.Task {
	var tasks []structs.Task
	db, _ := ConnectToSQLite()
	if show_deleted {
		if res := db.
			Find(&tasks); res.Error != nil {
			panic(res.Error)
		}
	} else {
		if res := db.
			Where("deleted = ?", 0).
			Find(&tasks); res.Error != nil {
			panic(res.Error)
		}
	}
	return tasks
}

func LogToFile(message string) {
	// Open the log file (create it if it doesn't exist)
	file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err) // If opening the file fails, log an error and exit
	}
	defer file.Close()

	// Set log output to the file
	logger := log.New(file, "", log.LstdFlags)
	logger.Println(message)
}
