package calendar

import (
	"fmt"
	"nug/helpers"
	"nug/structs"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func getTasks() []structs.Task {
	var tasks []structs.Task

	db, _ := helpers.ConnectToSQLite()

	if res := db.
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}

	return tasks
}

func (m Model) View(width, height int) string {
	month := time.Now().Local().Month()

	_, r := 2025, month

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("9")).
		MarginTop(1).
		Width(5).
		Align(lipgloss.Center)

	borderStyleToday := borderStyle.BorderForeground(lipgloss.Color("6"))
	borderStyleHeader := borderStyle.BorderForeground(lipgloss.Color("8"))
	borderStyleActive := borderStyle.BorderForeground(lipgloss.Color("3"))
	borderStyleTodayActive := borderStyle.BorderForeground(lipgloss.Color("0"))

	elements := []string{}
	firstWeekday := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	secondWeekday := time.Date(time.Now().Year(), time.Now().Month(), 2, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	thirdWeekday := time.Date(time.Now().Year(), time.Now().Month(), 3, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	fourthWeekday := time.Date(time.Now().Year(), time.Now().Month(), 4, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	fifthWeekday := time.Date(time.Now().Year(), time.Now().Month(), 5, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	sixthWeekday := time.Date(time.Now().Year(), time.Now().Month(), 6, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]
	seventWeekday := time.Date(time.Now().Year(), time.Now().Month(), 7, 0, 0, 0, 0, time.UTC).Local().Weekday().String()[0:3]

	left := ""
	left +=
		borderStyle.Width(width / 3).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyleHeader.Render(firstWeekday),
				borderStyleHeader.Render(secondWeekday),
				borderStyleHeader.Render(thirdWeekday),
				borderStyleHeader.Render(fourthWeekday),
				borderStyleHeader.Render(fifthWeekday),
				borderStyleHeader.Render(sixthWeekday),
				borderStyleHeader.Render(seventWeekday),
			),
		)

	for i := 1; i <= DaysInMonth(time.Now().Year(), r); i++ {
		row := ""
		if i == time.Now().Day() && i-1 == m.Selected {
			row += borderStyleTodayActive.Render(strconv.Itoa(i))
		} else if i == time.Now().Day() {
			row += borderStyleToday.Render(strconv.Itoa(i))
		} else if i-1 == m.Selected {
			row += borderStyleActive.Render(strconv.Itoa(i))
		} else {
			row += borderStyle.Render(strconv.Itoa(i))
		}

		elements = append(elements, row)
		if i%7 == 0 {
			left += lipgloss.JoinHorizontal(
				lipgloss.Top,
				elements...,
			)
			elements = []string{}
		}
	}

	right := ""
	for _, task := range getTasks() {

		year := time.Now().Year()
		month := time.Now().Month()

		selected_date := time.Date(year, month, m.Selected+1, 0, 0, 0, 0, time.UTC)

		if task.Time == selected_date.Format("02.01.2006") {
			right += fmt.Sprintf("Task: %s - %s \n", task.Name, task.Time)
		}
	}

	left = lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		borderStyle.Width(width/2).Render(
			right,
		),
	)

	return left
}

func DaysInMonth(year int, month time.Month) int {
	firstDayNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayCurrentMonth := firstDayNextMonth.AddDate(0, 0, -1)
	return lastDayCurrentMonth.Day()
}
