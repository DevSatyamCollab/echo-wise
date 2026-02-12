package main

import (
	"fmt"
	"log"

	internal "github.com/DevSatyamCollab/echo-wise/internal/core"
	"github.com/DevSatyamCollab/echo-wise/internal/suffle"
	"github.com/DevSatyamCollab/echo-wise/storage"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
)

const (
	defaultQuote  = "I’m not in this world to live up to your expectations and you’re not in this world to live up to mine."
	defaultAuthor = "Bruce Lee"
)

const (
	quoteInput = iota
	authorInput
)

// model
type model struct {
	inputs        []textinput.Model
	quotesList    []internal.Quote
	quote         string
	author        string
	style         styleBundle
	focused       int
	lastid        int
	store         *storage.Storage
	showInputForm bool
}

func InitialModel(s *storage.Storage) model {
	inputs := make([]textinput.Model, 2)
	// quote
	inputs[quoteInput] = textinput.New()
	inputs[quoteInput].Placeholder = "Don’t listen to what people say, watch what they do."
	inputs[quoteInput].Prompt = ""

	// author
	inputs[authorInput] = textinput.New()
	inputs[authorInput].Placeholder = "Churchill"
	inputs[authorInput].Prompt = ""

	// pre-define data
	ql := []internal.Quote{
		*internal.NewQuote(0,
			"Don't listen to what people say, watch what they do.",
			"Churchill",
		),

		*internal.NewQuote(1,
			"The only way to do great work is to love what you do.",
			"Steve Jobs",
		),

		*internal.NewQuote(2,
			"It is not the mountain we conquer, but ourselves.",
			"Sir Edmund Hilary",
		),

		*internal.NewQuote(3,
			"Success is not final, failure is not fatal: it is the courage to continue that counts.",
			"Winston Churchill",
		),

		*internal.NewQuote(4,
			"The trouble with having an open mind, of course, is that people will insist on coming along and trying to put things in it.",
			"Terry Pratchett",
		),

		*internal.NewQuote(5,
			"Life is what happens when you're busy making other plans.",
			"John Lennon",
		),

		*internal.NewQuote(6,
			"Everything is funny, as long as it's happening to somebody else",
			"Will Rogers",
		),

		*internal.NewQuote(7,
			"No act of kindness, no matter how small, is ever wasted.",
			"Aesop",
		),

		*internal.NewQuote(8,
			"In the middle or every difficult lies opportunity.",
			"Albert Einstein",
		),

		*internal.NewQuote(9,
			"Do what you can, with what you have, where you are.",
			"Theodore Roosevelt",
		),

		*internal.NewQuote(10,
			"Yesterday is history, tomorrow is a mystery, but today is a gift. That is why it is called the present",
			"Alice Morse Earle",
		),
	}

	// data from database
	list, err := s.GetData()
	if err != nil {
		log.Fatalf("Error can't get the data from database: %v", err)
	}

	// only one time
	if len(list) < 10 {
		// insert some pre-defined database
		for _, q := range ql {
			if err := s.AddData(q.Quote, q.Author); err != nil {
				log.Printf("Error can't add data to the database: %v\n", err)
			}
		}

		list = ql

	}

	return model{
		style:      DefaultStyle(65),
		store:      s,
		inputs:     inputs,
		quotesList: list,
		quote:      defaultQuote,
		author:     defaultAuthor,
	}
}

// validate inputs

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width < 65 {
			m.style = DefaultStyle(msg.Width)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "r":
			if !m.showInputForm {
				for {
					q := suffle.Suffle(m.quotesList)
					if q.Id != m.lastid {
						m.lastid = q.Id
						m.quote = q.Quote
						m.author = q.Author
						break
					}
				}
			}
		case "a":
			if !m.showInputForm {
				m.showInputForm = true
				m.focused = quoteInput
				m.inputs[authorInput].Blur()
				m.inputs[m.focused].Focus()
				return m, nil
			}
		case "ctrl+l":
			if !m.showInputForm {

			}
		case "esc":
			m.backToMain()
		case "enter":
			if m.showInputForm {
				if m.focused == len(m.inputs)-1 {
					if m.inputs[quoteInput].Value() != "" {
						m.quote = m.inputs[quoteInput].Value()
						m.author = m.inputs[authorInput].Value()

						// add to the current list
						q := *internal.NewQuote(len(m.quotesList),
							m.inputs[quoteInput].Value(),
							m.inputs[authorInput].Value())
						m.quotesList = append(m.quotesList, q)

						// save to the database
						go func() {
							if err := m.store.AddData(q.Quote, q.Author); err != nil {
								log.Printf("Error can' save to the database: %v", err)
							}
						}()
						// reset
						m.backToMain()
						return m, nil
					}
				}
				m.nextInput()
			}
		case "shift+tab":
			if m.showInputForm {
				m.prevInput()
			}

		case "tab":
			if m.showInputForm {
				m.nextInput()
			}
		}

		if m.showInputForm {
			for i := range m.inputs {
				m.inputs[i].Blur()
			}

			m.inputs[m.focused].Focus()
		}

	}

	if m.showInputForm {
		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m model) View() string {
	var footer, header, mainContent, innerContent, ui, qtext, atext string

	// main View
	header = m.style.header.Render("Quote of the day")
	footer = m.style.footer.Render("r: reload . a: add a Quote . l: list of Quote . q: Quit")
	qtext = m.style.quoteText.Render(fmt.Sprintf("\"%s\"", m.quote))
	atext = m.style.author.Render("--" + m.author)

	// join the vertically
	innerContent = lipgloss.JoinVertical(lipgloss.Left, qtext, atext)
	mainContent = m.style.quoteBox.Render(innerContent)

	// add view
	if m.showInputForm {
		footer = m.style.footer.Render("esc: back . tab: next . shift+tab: prev . enter: continue")
		formFields := lipgloss.JoinVertical(
			lipgloss.Left,
			inputStyle.Render("Quote..."),
			m.inputs[quoteInput].View(),
			"\n", // little space
			inputStyle.Render("Author"),
			m.inputs[authorInput].View(),
			"\n",
		)

		mainContent = m.style.quoteBox.Render(formFields)
	}

	ui = lipgloss.JoinVertical(lipgloss.Left, header, mainContent, footer)
	return m.style.container.Render(ui)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func (m *model) backToMain() {
	m.inputs[quoteInput].SetValue("")
	m.inputs[authorInput].SetValue("")
	m.showInputForm = false
	m.focused = 0
}
