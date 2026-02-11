package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	continueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#767676"))
)

const (
	quoteInput = iota
	authorInput
)

// model
type model struct {
	style         styleBundle
	inputs        []textinput.Model
	focused       int
	quote         string
	author        string
	showInputForm bool
}

func InitialModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	// quote
	inputs[quoteInput] = textinput.New()
	inputs[quoteInput].Placeholder = "Don’t listen to what people say, watch what they do."
	inputs[quoteInput].Prompt = ""
	//inputs[quoteInput].Validate =

	// author
	inputs[authorInput] = textinput.New()
	inputs[authorInput].Placeholder = "Churchill"
	inputs[authorInput].Prompt = ""

	return model{
		style:  DefaultStyle(65),
		inputs: inputs,
		quote:  "Quote",
		author: "Author",
	}
}

// validate inputs

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

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
		case "a":
			if !m.showInputForm {
				m.showInputForm = true
				m.focused = quoteInput
				m.inputs[authorInput].Blur()
				m.inputs[m.focused].Focus()
				return m, nil
			}
		case "esc":
			m.backToMain()
		case "enter":
			if m.showInputForm {
				if m.focused == len(m.inputs)-1 {
					if m.inputs[quoteInput].Value() != "" {
						m.quote = m.inputs[quoteInput].Value()
						m.author = m.inputs[authorInput].Value()

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
	if m.quote == "" {
		m.quote = "Quote"
	}

	if m.author == "" {
		m.author = ""
	}

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
