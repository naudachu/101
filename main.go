package main

import (
	"101/player"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/color"
)

//[ ] https://dev.to/pomdtr/how-to-debug-bubble-tea-applications-in-visual-studio-code-50jp

// Figure out how to connect to headless dlv

var players []*player.Player

type model struct {
	list     []*player.Player
	cursor   int
	selected int
	chosen   bool
	command  string

	textInput textinput.Model
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initModel() model {
	flag.Parse()
	names := flag.Args()

	switch len(names) {
	case 0:
		players = append(players, player.NewPlayer("Player 1"))
		players = append(players, player.NewPlayer("Player 2"))
	default:
		for _, e := range names {
			{
				players = append(players, player.NewPlayer(e))
			}
		}
	}

	return model{
		list:   players,
		cursor: 0,
	}
}

func (m *model) initInputField() {
	ti := textinput.New()
	ti.Placeholder = "Command for..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	m.textInput = ti
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.chosen {
		return playersUpdate(msg, m)
	} else {
		return scoreUpdate(msg, m)
	}
}

func scoreUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:

			m.convertStringToCommand(m.textInput.Value())
			m.chosen = false
			m.cursor = 0
			return m, nil
		}

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}
	return m, nil
}

func playersUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.list)-1 {
				m.cursor++
			}

		case "enter":
			m.chosen = true
			m.selected = m.cursor
			m.initInputField()
			m.command = "add"

		case "-":
			m.chosen = true
			m.selected = m.cursor
			m.initInputField()
			m.command = "sub"
		}

	}

	return m, nil
}

func (m model) View() string {
	var result string
	if !m.chosen {
		result = playersView(m)
	} else {
		result = scoreView(m)
	}
	return result
}

func scoreView(m model) string {
	return fmt.Sprintf(
		"Add round calculation command:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func playersView(m model) string {
	var output string
	for i, p := range m.list {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		output += cursor
		output += color.HEX(p.Color()).Sprintf("%d. %s\n", i, p.Name())
		output += fmt.Sprintln(p.Score())
	}

	return output

}

func (m *model) convertStringToCommand(text string) {
	switch m.command {
	case "add":

		players[m.selected].SetPoints(text)

	case "sub":

		// Check if -points can be converted to int
		subPoints, err := strconv.Atoi(text)
		if err != nil {
			log.Println("subPoints convertion to int failed")
		}

		players[m.selected].SubPoints(subPoints)

	default:
		log.Println("wrong command")
	}
}
