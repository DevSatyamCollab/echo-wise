package main

import "github.com/charmbracelet/lipgloss"

// style
type styleBundle struct {
	header lipgloss.Style
	footer lipgloss.Style
	quote  lipgloss.Style
	author lipgloss.Style
}

func DefaultStyle() styleBundle {
	return styleBundle{
		header: lipgloss.NewStyle().Foreground(lipgloss.Color("201")),
		footer: lipgloss.NewStyle().Foreground(lipgloss.Color("196")),
		quote: lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("82")).
			Padding(1, 2),
	}
}
