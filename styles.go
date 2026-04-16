package main

import "github.com/charmbracelet/lipgloss"

const (
	colorPurple = lipgloss.Color("#b298ff")
	colorGray   = lipgloss.Color("#7F7F7F")
	colorWhite  = lipgloss.Color("#FFFFFF")
	colorBlack  = lipgloss.Color("#000000")
	colorDark   = lipgloss.Color("#1A1A1A")
	colorRed    = lipgloss.Color("#FF6B6B")
)

var (
	topBarStyle = lipgloss.NewStyle().
			Background(colorDark).
			Foreground(colorGray).
			Padding(0, 1)

	topBarLabelStyle = lipgloss.NewStyle().
				Foreground(colorPurple).
				Bold(true)

	topBarDimStyle = lipgloss.NewStyle().
			Foreground(colorGray)
)

var (
	menuItemStyle = lipgloss.NewStyle().
			Foreground(colorWhite).
			Padding(0, 2)

	menuSelectedStyle = lipgloss.NewStyle().
				Foreground(colorBlack).
				Background(colorPurple).
				Padding(0, 2).
				Bold(true)

	menuDescStyle = lipgloss.NewStyle().
			Foreground(colorGray).
			Padding(0, 4)
)

var (
	tableHeaderStyle = lipgloss.NewStyle().
				Foreground(colorPurple).
				Bold(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(colorGray)

	cellStyle = lipgloss.NewStyle().
			Foreground(colorWhite)

	cellMutedStyle = lipgloss.NewStyle().
			Foreground(colorGray)
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(colorPurple).
			Bold(true)

	badgeStyle = lipgloss.NewStyle().
			Foreground(colorBlack).
			Background(colorPurple).
			Padding(0, 1).
			Bold(true)

	aiStyle = lipgloss.NewStyle().
		Foreground(colorPurple).
		Bold(true)

	userStyle = lipgloss.NewStyle().
			Foreground(colorGray).
			Bold(true)

	contentStyle = lipgloss.NewStyle().
			Foreground(colorWhite)
)

var (
	dimStyle = lipgloss.NewStyle().
			Foreground(colorGray)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorRed).
			Bold(true)
)

const catASCII = `
    ██████                      ██████
    ██░░██████              ██████░░██
    ██░░░░░░████  ██████  ████░░░░░░██
    ██░░░░░░░░██████████████░░░░░░░░██
    ████░░██████████████████████░░████
      ██████████████████████████████
        ██████████████████████████
      ██████████████████████████████
      ██████    ██████████    ██████
      ████    ██  ██████  ██    ████
██████████    ██  ██████  ██    ██████████
        ████    ██████████    ████
      ██████████████░░██████████████
    ████    ██████████████████    ████

	 `
