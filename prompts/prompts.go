//go:build !js
// +build !js

package prompts

import (
	"fmt"
	"os"
	"strings"

	"github.com/BryanMwangi/go-agent/config"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input  textinput.Model
	list   list.Model
	choice string
	mode   string // "username", "apikey", "model", etc.
	items  []string
	done   bool
}

func initialModel(label string, mode string, items []string) model {
	ti := textinput.New()
	ti.Placeholder = label
	ti.Focus()
	ti.CharLimit = 500
	ti.Width = 40

	var l list.Model
	if len(items) > 0 {
		emptyList := []list.Item{}
		l = list.New(emptyList, list.NewDefaultDelegate(), 40, 10)
		l.Title = label
		l.SetItems(itemsToListItems(items))
	}

	return model{
		input: ti,
		list:  l,
		mode:  mode,
		items: items,
	}
}

func itemsToListItems(items []string) []list.Item {
	var result []list.Item
	for _, i := range items {
		result = append(result, listItem(i))
	}
	return result
}

type listItem string

func (i listItem) Title() string       { return string(i) }
func (i listItem) Description() string { return "" }
func (i listItem) FilterValue() string { return string(i) }

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q":
			return m, tea.Quit
		case "ctrl+c":
			os.Exit(0)
			return m, tea.Quit
		case "enter":
			if m.mode == "model" && len(m.items) > 0 {
				m.choice = m.list.SelectedItem().(listItem).Title()
				m.done = true
				return m, tea.Quit
			} else {
				inputVal := m.input.Value()
				m.choice = strings.TrimSpace(inputVal)
				m.done = true
				return m, tea.Quit
			}
		}
	}

	if m.mode == "model" && len(m.items) > 0 {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	} else {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.done {
		return ""
	}
	if m.mode == "model" && len(m.items) > 0 {
		return m.list.View()
	}
	return fmt.Sprintf("%s\n\n%s", m.input.Placeholder, m.input.View())
}

func promptBubbletea(label, mode string, items []string) string {
	m := initialModel(label, mode, items)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error running prompt:", err)
		os.Exit(1)
	}
	return finalModel.(model).choice
}

func PromptUsername() string {
	return promptBubbletea("Enter your name:", "username", nil)
}

func PromptAPIKey() string {
	return promptBubbletea("Enter your OpenAI API key:", "apikey", nil)
}

func PromptModel() string {
	return promptBubbletea("Select a model:", "model", config.AvailableModels)
}

func PromptWorkingDirectory() string {
	return promptBubbletea("Enter your working directory:", "workdir", nil)
}

func PromptUserInput() string {
	return promptBubbletea("Start a new conversation:", "userinput", nil)
}
