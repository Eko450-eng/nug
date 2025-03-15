package mainapp

func (m model) View() string {
	s := ""

	switch m.state {
	case mainState:
		s += m.taskoverview.View(m.width, m.height)

	case helpState:
		s = m.helpmodal.View(m.width, m.height)
	}

	return s
}
