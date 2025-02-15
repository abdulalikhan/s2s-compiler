package constants

import "github.com/charmbracelet/lipgloss"

var (
	RedText    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	GreenText  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00"))
	YellowText = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFF"))
)

// helper functions to apply the style
func StyledYellow(text string) string {
	return YellowText.Render(text)
}

func StyledGreen(text string) string {
	return GreenText.Render(text)
}

func StyledRed(text string) string {
	return RedText.Render(text)
}
