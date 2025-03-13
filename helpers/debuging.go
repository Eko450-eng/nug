package helpers

import (
	"fmt"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		LogToFile(fmt.Sprintf("%s", err))
	}
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
