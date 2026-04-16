package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	state  viewState
	cursor int
	width  int

	viewport viewport.Model
	textarea textarea.Model
	spinner  spinner.Model

	messages       []chatMessage
	pendingCommand string

	cats       []Cat
	categories []CatCategory
	totalItems int
	totalPages int

	loading bool
	err     error
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Ask me anything..."
	ta.Focus()
	ta.Prompt = ""
	ta.CharLimit = 280
	ta.SetWidth(80)
	ta.SetHeight(1)
	ta.ShowLineNumbers = false
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	vp := viewport.New(80, 10)
	vp.SetContent(renderWelcome())

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(colorPurple)

	return model{
		state:    viewMenu,
		viewport: vp,
		textarea: ta,
		spinner:  sp,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}
