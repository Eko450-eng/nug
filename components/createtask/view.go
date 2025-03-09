package createtask

func (m CreateModel) View() string {
	current := m.Fields[m.EditLine]
	return m.styles.InputField.Render(current.InputField.View())
}
