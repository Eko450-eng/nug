package calendar

import (
	"fmt"
	"nug/helpers"
	"nug/structs"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func dayHasTask(date string) []structs.Task {
	var tasks []structs.Task

	db, _ := helpers.ConnectToSQLite()

	if res := db.
		Where("date = ?", helpers.NormalizeDate(date)).
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}

	return tasks
}

func getTasks() []structs.Task {
	var tasks []structs.Task

	db, _ := helpers.ConnectToSQLite()

	if res := db.
		Find(&tasks); res.Error != nil {
		panic(res.Error)
	}

	return tasks
}

var borderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("9")).
	MarginTop(1).
	Width(50 / 7).
	Align(lipgloss.Center)

var borderStyleHeader = borderStyle.BorderForeground(lipgloss.Color("#8ecae6"))
var borderStyleToday = borderStyle.BorderForeground(lipgloss.Color("#219ebc"))
var borderStyleActive = borderStyle.BorderForeground(lipgloss.Color("#ffb703"))
var borderStyleTodayActive = borderStyle.BorderForeground(lipgloss.Color("#c1121f"))

func displayHeader() string {
	res := fmt.Sprintf("%s", time.Now().Month().String())

	res += lipgloss.JoinHorizontal(
		lipgloss.Top,
		borderStyleHeader.Render("Mon"),
		borderStyleHeader.Render("Tue"),
		borderStyleHeader.Render("Wed"),
		borderStyleHeader.Render("Thu"),
		borderStyleHeader.Render("Fri"),
		borderStyleHeader.Render("Sat"),
		borderStyleHeader.Render("Sun"),
	)

	return res
}

func (m Model) displayWeekLine() string {
	res := ""

	month := time.Now().Local().Month()

	_, r := 2025, month

	elements := []string{}

	placeholders := 0
	firstDay := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC).Local().Weekday()
	switch firstDay {
	case time.Monday:
		placeholders += 0
	case time.Tuesday:
		placeholders += 1
	case time.Wednesday:
		placeholders += 2
	case time.Thursday:
		placeholders += 3
	case time.Friday:
		placeholders += 4
	case time.Saturday:
		placeholders += 5
	case time.Sunday:
		placeholders += 6
	}

	placeHolderRows := ""
	style := borderStyle.BorderForeground(lipgloss.Color("#c1121f"))

	boxes := []string{}

	for p := 1; p <= placeholders; p++ {
		boxes = append(boxes, style.Render(" "))
	}

	placeHolderRows += lipgloss.JoinHorizontal(
		lipgloss.Top,
		boxes...,
	)

	placeHolderRowsJoined := lipgloss.JoinHorizontal(
		lipgloss.Top,
		placeHolderRows,
	)

	elements = append(elements, placeHolderRowsJoined)

	for i := 1; i <= DaysInMonth(time.Now().Year(), r); i++ {
		row := ""

		if time.Now().Day() == i && i-1 == m.Selected {
			style = borderStyleTodayActive
		} else if i == time.Now().Day() {
			style = borderStyleToday
		} else if i-1 == m.Selected {
			style = borderStyleActive
		} else {
			style = borderStyle
		}

		year := strconv.Itoa(time.Now().Year())
		month := time.Now().Month()

		tasks := dayHasTask(fmt.Sprintf("%s.%s.%s", strconv.Itoa(i), strconv.Itoa(int(month)), year))

		if len(tasks) > 0 {
			row += style.Foreground(lipgloss.Color("#c1121f")).Render(strconv.Itoa(i))
		} else {
			row += style.Render(strconv.Itoa(i))
		}

		elements = append(elements, row)
		if (i+placeholders)%7 == 0 {
			res += lipgloss.JoinHorizontal(
				lipgloss.Top,
				elements...,
			)
			elements = []string{}
		}
	}

	if len(elements) > 0 {
		res += lipgloss.JoinHorizontal(
			lipgloss.Top,
			elements...,
		)
	}

	return res
}

func (m Model) View(width, height int) string {
	right := ""
	for _, task := range getTasks() {

		year := time.Now().Year()
		month := time.Now().Month()

		selected_date := time.Date(year, month, m.Selected+1, 0, 0, 0, 0, time.UTC)

		if task.Date == selected_date.Format("2.1.2006") {
			right += fmt.Sprintf("Task: %s - %s \n", task.Name, task.Date)
		}
	}

	left := lipgloss.JoinVertical(
		lipgloss.Top,
		displayHeader(),
		m.displayWeekLine(),
	)

	res := lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		borderStyle.
			Width(width/2).
			Height(height/2).
			Render(right),
	)

	return res
}

func DaysInMonth(year int, month time.Month) int {
	firstDayNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayCurrentMonth := firstDayNextMonth.AddDate(0, 0, -1)
	return lastDayCurrentMonth.Day()
}
