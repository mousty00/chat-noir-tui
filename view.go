package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	topBar := renderTopBar(m)
	vp := m.viewport.View()

	if m.state == viewChat {
		prompt := lipgloss.JoinHorizontal(lipgloss.Bottom,
			" ",
			badgeStyle.Render(" Chat Noir "),
			" ",
			m.textarea.View(),
		)
		return lipgloss.JoinVertical(lipgloss.Top, topBar, " ", vp, prompt)
	}

	return lipgloss.JoinVertical(lipgloss.Top, topBar, " ", vp)
}

func renderTopBar(m model) string {
	label := m.state.String()
	padding := max(0, m.width-len(" Chat Noir ")-len(" / "+label+" ")-12)

	return topBarStyle.Width(m.width).Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			topBarLabelStyle.Render(" Chat Noir "),
			topBarDimStyle.Render(" / "+label+" "),
			strings.Repeat(" ", padding),
			topBarDimStyle.Render(" Esc back "),
		),
	)
}
