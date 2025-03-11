package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"nug/structs"

	"net/http"

	"github.com/glebarez/sqlite"
	gap "github.com/muesli/go-app-paths"
	"gorm.io/gorm"
)

func SyncToWebDav() {
	url := os.Getenv("NUG_URL")
	username := os.Getenv("NUG_USERNAME")
	password := os.Getenv("NUG_PASSWORD")

	LogToFile(fmt.Sprintf("URL: %s", url))
	LogToFile(fmt.Sprintf("User: %s", username))
	LogToFile(fmt.Sprintf("pass: %s", password))

	scope := gap.NewScope(gap.User, "nug")
	dirs, err := scope.DataDirs()
	appPath := dirs[0] + "/nug.db"

	file, err := os.Open(appPath)
	req, err := http.NewRequest("PUT", url, file)

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	req.Header.Set("Overwrite", "T")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	CheckErr(err)

	// Check for success
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		LogToFile(fmt.Sprintf("failed to upload, status: %s", resp.Status))
	}
}

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

func UpdateTasks() []structs.Task {
	var tasks []structs.Task
	db, _ := ConnectToSQLite()
	// if show_deleted {
	// 	if res := db.
	// 		Find(&tasks); res.Error != nil {
	// 		panic(res.Error)
	// 	}
	// } else {
	if res := db.
		Where("deleted = ?", 0).
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}
	// }
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

func NormalizeDate(date string) string {
	parts := strings.Split(date, ".")
	if len(parts) != 3 {
		return date // Invalid format, return as-is
	}

	day, _ := strconv.Atoi(parts[0])   // Convert to int to remove leading zeros
	month, _ := strconv.Atoi(parts[1]) // Convert to int to remove leading zeros
	year := parts[2]                   // Keep year as a string

	return fmt.Sprintf("%d.%d.%s", day, month, year)
}

func ShortenString(str string, length int) string {
	if len(str) > length {
		return str[:15]
	}
	return str
}
