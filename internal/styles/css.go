package internal

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	AppStyle      lipgloss.Style
	TitleDefault  lipgloss.Style
	TitleSelected lipgloss.Style
	DescDefault   lipgloss.Style
	DescSelected  lipgloss.Style
	BlockDefault  lipgloss.Style
	BlockSelected lipgloss.Style
	BlockHelp     lipgloss.Style
	HelpKeyStyle  lipgloss.Style
	HelpDescStyle lipgloss.Style
}

func New() *Styles {
	var s Styles
	s.AppStyle = lipgloss.NewStyle().Margin(1, 2)

	s.TitleDefault = lipgloss.NewStyle()

	s.TitleSelected = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

	s.DescDefault = s.TitleDefault.
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#C3C3C3"})

	s.DescSelected = s.TitleSelected.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	s.BlockDefault = lipgloss.NewStyle().
		Height(2).
		Padding(0, 0, 0, 1).
		Margin(1, 0, 1, 1)

	s.BlockSelected = s.BlockDefault.Copy().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Margin(1, 0, 1, 0)

	s.BlockHelp = lipgloss.NewStyle().
		Height(2).
		Padding(0, 0, 0, 0).
		Margin(3, 0, 1, 1)

	s.HelpKeyStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	})

	s.HelpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#B2B2B2",
		Dark:  "#4A4A4A",
	})
	return &s
}
