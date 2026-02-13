package main

import (
	core "github.com/DevSatyamCollab/echo-wise/internal/core"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	quote, author string
	id            int
}

func (i item) Title() string       { return i.quote }
func (i item) Description() string { return i.author }
func (i item) FilterValue() string { return i.author }

func ListQuotesItems(qlist []core.Quote) []list.Item {
	items := make([]list.Item, len(qlist))
	for i, q := range qlist {
		items[i] = item{id: q.Id, quote: q.Quote, author: q.Author}
	}

	return items
}

/*
func ItemsToQuotes(items []item) []core.Quote {
	ql := make([]core.Quote,len(items))
	for _, item := range items {
		ql = append(ql, )
	}
	return ql
}
*/
