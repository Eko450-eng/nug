package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"nask/structs"

	"github.com/glebarez/sqlite"
	// gap "github.com/muesli/go-app-paths"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	// scope := gap.NewScope(gap.User, "nask")
	// dirs, err := scope.DataDirs()
	// app_path := dirs[0] + "nask.db"
	// CheckErr(err)
	// app_path := "/home/eko/.local/share/nask/nask.db"
	app_path := "./nask.db"

	_, e := os.Stat(app_path)
	CheckErr(e)

	if os.IsNotExist(e) {
		os.MkdirAll(app_path, 0o770)
	}

	db, err := gorm.Open(sqlite.Open(app_path), &gorm.Config{})
	CheckErr(err)

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
