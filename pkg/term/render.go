package term

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

func Serve(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	return initModel(), []tea.ProgramOption{
		tea.WithInput(sess),
		tea.WithOutput(sess),
	}
}

func Render() tea.Model {
	return initModel()
}

type model struct {
	toggle bool
}

func initModel() model {
	return model{toggle: false}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.toggle = !m.toggle
		}

	}
	return m, nil
}

func (m model) View() string {

	if m.toggle == true {
		return yay
	}

	return greet
}
