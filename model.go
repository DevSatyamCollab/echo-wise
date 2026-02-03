package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// model
type model struct {
}

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	header := "Quote of the day"
	footer := "r: reload . a: add a quote . q: Quit . h: help"
	view := ""

	return fmt.Sprintf("\n%s\n\n%s\n\n", header, view, footer)
}
