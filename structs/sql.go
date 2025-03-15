package structs

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string
	Description string
	Project_id  int
	Prio        int
	Time        string
	Date        string
	Deletedtime string
	Modified    string
	Completed   int
	Deleted     int
}

type Tags struct {
	gorm.Model
	Name        string
	Deleted     int
	Deletedtime string
}

type Tag_to_task struct {
	gorm.Model
	Tag  int
	Task int
}

type Project struct {
	gorm.Model
	Name        string
	Deleted     int
	Deletedtime string
}

type Task_to_Project struct {
	gorm.Model
	Task    int
	Project int
}
