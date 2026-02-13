package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"time"

	predefineddata "github.com/DevSatyamCollab/echo-wise/internal/PreDefinedData"
	core "github.com/DevSatyamCollab/echo-wise/internal/core"
	"github.com/DevSatyamCollab/echo-wise/internal/suffle"
	"github.com/DevSatyamCollab/echo-wise/storage"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	defaultWidth = 65
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
	quotesList    []core.Quote
	list          list.Model
	spinner       spinner.Model
	timer         timer.Model
	quote         string
	author        string
	style         styleBundle
	focused       int
	lastid        int
	store         *storage.Storage
	showInputForm bool
	loading       bool
	showingList   bool
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

	// data from database
	qlist, err := s.GetData()
	if err != nil {
		log.Fatalf("Error can't get the data from database: %v", err)
	}

	// app set up (one time only)
	if len(qlist) < 10 {
		// insert some pre-defined database
		ql := predefineddata.GetPreData()
		for _, q := range ql {
			if err := s.AddData(q.Quote, q.Author); err != nil {
				log.Printf("Error can't add data to the database: %v\n", err)
			}
		}

		qlist = ql
	}

	// spinner setup
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	// list setup
	l := list.New(ListQuotesItems(qlist), list.NewDefaultDelegate(), 0, 0)
	l.Title = "All Quotes"

	// add a delete key in list
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("ctrl+d"),
				key.WithHelp("ctrl+d", "delete quote"),
			),
			key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back to main menu"),
			),
		}
	}

	return model{
		style:      DefaultStyle(65),
		store:      s,
		inputs:     inputs,
		quotesList: qlist,
		quote:      defaultQuote,
		author:     defaultAuthor,
		spinner:    sp,
		list:       l,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleResize(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// quit the program
		case "q":
			return m, tea.Quit

		// reload the quote
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

				m.loading = true
				m.timer = timer.New(100 * time.Millisecond)
				return m, tea.Batch(m.spinner.Tick, m.timer.Init())
			}

		// add a new quote
		case "a":
			if !m.showInputForm {
				m.showInputForm = true
				m.focused = quoteInput

				return m, m.inputs[m.focused].Focus()
			}

		case "ctrl+d":
			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					delIndex, found := slices.BinarySearchFunc(m.quotesList, item.id, func(q core.Quote, target int) int {
						return cmp.Compare(q.Id, target)
					})

					if found {
						m.quotesList = slices.Delete(m.quotesList, delIndex, delIndex+1)
						m.list.SetItems(ListQuotesItems(m.quotesList))
						if err := m.store.DeleteData(item.id); err != nil {
							log.Fatalln(err)
						}
					}
				}
			}
		// show the list of quotes
		case "ctrl+l":
			if !m.showInputForm && !m.loading {
				m.showingList = true
				m.list.SetItems(ListQuotesItems(m.quotesList))
				return m, nil
			}

		// back to main menu
		case "esc":
			m.backToMain()

		// continue
		case "enter":
			if m.showInputForm {
				if m.focused == len(m.inputs)-1 {
					if m.inputs[quoteInput].Value() != "" {
						m.quote = m.inputs[quoteInput].Value()
						m.author = m.inputs[authorInput].Value()

						// add to the current list
						q := *core.NewQuote(len(m.quotesList),
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

	// Tigger when time reaches 0
	case timer.TimeoutMsg:
		m.loading = false
		return m, nil

	// Necessary to keep timer's internal state ticking
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	// Necessary to keep spinner animating
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.showInputForm {
		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
			m.inputs[i].Blur()
			m.inputs[m.focused].Focus()
		}

		return m, tea.Batch(cmds...)
	}

	if m.showingList {
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	var footer, header, mainContent, innerContent, ui, qtext, atext string

	// main View
	header = m.style.header.Render("Quote of the day")
	footer = m.style.footer.Render("r: reload . a: add a Quote . l: list of Quote . q: Quit")
	qtext = m.style.quoteText.Render(fmt.Sprintf("\"%s\"", m.quote))
	atext = m.style.author.Render("― " + m.author)

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

	// spinner view
	if m.loading {
		mainContent = fmt.Sprintf("\n %s Loading...\n", m.spinner.View())
	}

	// list view
	if m.showingList {
		return docStyle.Render(m.list.View())
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
	m.showingList = false
}

func (m *model) handleResize(msg tea.WindowSizeMsg) {
	_, v := docStyle.GetFrameSize()
	width := min(msg.Width, defaultWidth)
	m.list.SetSize(width, msg.Height-v) // ← Use width directly, only subtract v
	m.style = DefaultStyle(width)
}
