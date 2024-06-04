package term

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

func Serve(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	return initMenu(), []tea.ProgramOption{
		tea.WithInput(sess),
		tea.WithOutput(sess),
	}
}

func Render() tea.Model {
	return initMenu()
}

type menu struct {
	opts   []string
	cursor int
}

type image struct {
	content string
}

func initMenu() menu {
	return menu{
		opts:   []string{"yay", "greet"},
		cursor: 0,
	}
}

func initImag(str string) image {
	return image{
		str,
	}
}

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j":
			if m.cursor == len(m.opts)-1 {
				m.cursor = 0
			} else {
				m.cursor += 1
			}
		case "k":
			if m.cursor == 0 {
				m.cursor = len(m.opts) - 1
			} else {
				m.cursor -= 1
			}
		case "enter":
			return initImag(ascii[m.opts[m.cursor]]), nil
		}

	}
	return m, nil
}

func (m menu) View() string {
	s := "choose:\n\n"
	for i, choice := range m.opts {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m image) Init() tea.Cmd {
	return nil
}

func (m image) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return initMenu(), nil
		}
	}

	return m, nil
}

func (m image) View() string {
	return m.content
}
