package main

import (
	"101/internal/player"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	css "101/internal/styles"
)

//[ ] https://dev.to/pomdtr/how-to-debug-bubble-tea-applications-in-visual-studio-code-50jp

// Figure out how to connect to headless dlv

var (
	players []*player.Player
)

type model struct {
	s        css.Styles
	list     []*player.Player
	cursor   int
	selected int
	chosen   bool
	command  string

	textInput textinput.Model
	listHelp  string
}

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initModel() model {
	flag.Parse()
	names := flag.Args()

	{ // Create players from CLI args
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
	}

	csss := *css.New()

	return model{
		list:     players,
		cursor:   0,
		s:        csss,
		listHelp: initListHelp(&csss),
	}
}

func initListHelp(s *css.Styles) string {
	help := map[string]string{
		"enter": "add new points",
		"bckp":  "substract int",
		"0":     "nullify player",
		"esc":   "quit",
	}

	var helpWriter strings.Builder
	for i, j := range help {
		helpWriter.WriteString(fmt.Sprintf(
			"%s %s\n",
			s.HelpKeyStyle.Render("• "+i),
			s.HelpDescStyle.Render(j)))
	}
	return helpWriter.String()
}

func (m *model) initInputField() {
	ti := textinput.New()
	switch m.command {
	case "add":
		ti.Placeholder = "jqk67890t"
	case "sub":
		ti.Placeholder = "qQ / qqQQ"
	}
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

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
		switch msg.String() {
		case "j", "J", "k", "K", "q", "Q", "t", "T", "6", "7", "8", "9", "0":
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func playersUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "up":
			m.CursorUp()

		case "down":
			m.CursorDown()

		case "enter":
			m.SelectWithAddCmd()
		case "backspace":
			m.SelectWithSubstractCmd()
		case "0":
			m.newFunction()
		// These keys should exit the program.
		case "esc":
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m *model) newFunction() {
	m.selected = m.cursor
	m.list[m.selected].Win()
}

func (m *model) CursorUp() {
	if m.cursor > 0 {
		m.cursor--
	} else {
		m.cursor = len(m.list) - 1
	}
}

func (m *model) CursorDown() {
	if m.cursor < len(m.list)-1 {
		m.cursor++
	} else {
		m.cursor = 0
	}
}

func (m *model) SelectWithAddCmd() {
	m.chosen = true
	m.selected = m.cursor
	m.command = "add"
	m.initInputField()
}

func (m *model) SelectWithSubstractCmd() {
	m.chosen = true
	m.selected = m.cursor
	m.command = "sub"
	m.initInputField()
}

func (m model) View() string {
	var rndr string
	if !m.chosen {
		rndr = playersView(m)
	} else {
		rndr = scoreView(m)
	}
	return rndr
}

func scoreView(m model) string {
	var prompt string
	switch m.command {
	case "add":
		prompt = "Card combination"
	case "sub":
		prompt = "Q or QQ"
	}

	prompt = m.s.TitleSelected.Render(prompt)
	help := m.s.HelpDescStyle.Render("esc -> quit")

	return m.s.AppStyle.Render(fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		prompt,
		m.textInput.View(),
		help,
	) + "\n")
}

func playersView(m model) string {
	var sb strings.Builder

	for i, p := range m.list {
		switch m.cursor {
		case i: // selected
			title := m.s.TitleSelected.Render(p.Title())
			desc := m.s.DescSelected.Render(p.Description())
			content := fmt.Sprintf(
				"%s\n%s",
				title,
				desc)

			sb.WriteString(m.s.BlockSelected.Render(content))
		default:
			title := m.s.TitleDefault.Render(p.Title())
			desc := m.s.DescDefault.Render(p.Description())
			content := fmt.Sprintf(
				"%s\n%s",
				title,
				desc)

			sb.WriteString(m.s.BlockDefault.Render(content))
		}
	}

	sb.WriteString(m.s.BlockHelp.Render(m.listHelp))

	return m.s.AppStyle.Render(sb.String())
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
