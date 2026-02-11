package main

import "github.com/charmbracelet/lipgloss"

// style
type styleBundle struct {
	header    lipgloss.Style
	footer    lipgloss.Style
	container lipgloss.Style // global outer border (program border)
	quoteBox  lipgloss.Style // inner border
	author    lipgloss.Style
	quoteText lipgloss.Style
}

func DefaultStyle(width int) styleBundle {
	const padding = 2
	// make it flexible to windows size
	contentWidth := width - 4

	return styleBundle{
		header: lipgloss.NewStyle().
			Foreground(lipgloss.Color("201")).
			MarginLeft(1).MarginBottom(1),

		footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			MarginLeft(1).MarginTop(1),

		container: lipgloss.NewStyle().
			Padding(1).
			Width(contentWidth),

		quoteBox: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("82")).
			Padding(1, 2).
			Width(contentWidth - (padding * 2)),

		quoteText: lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Align(lipgloss.Center).
			Width(contentWidth - 8),

		author: lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Align(lipgloss.Right).
			Width(contentWidth - 8),
	}
}
