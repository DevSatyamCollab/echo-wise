package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// model
type model struct {
	style styleBundle
}

func InitialModel() model {
	return model{
		style: DefaultStyle(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "r":
		case "a":
		}
	}

	return m, nil
}

func (m model) View() string {
	header := m.style.header.Render("Quote of the day")
	footer := m.style.footer.Render("r: reload . a: add a quote . q: Quit")
	view := m.style.quote.Render("Quote")

	return fmt.Sprintf("\n%s\n\n%s\n\n%s\n", header, view, footer)
}
