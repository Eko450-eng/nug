package taskoverview

import (
	"fmt"
)

func (m Model) View() string {
	tasks := ""

	for i, task := range m.Tasks {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.Tasks[i].Completed == 1 {
			checked = "x"
		}
		if m.Tasks[i].Deleted == 1 {
			tasks += fmt.Sprintf("%s D-[%s] %s\n", cursor, checked, task.Name)
		} else {
			tasks += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Name)
		}
	}

	return tasks
}
