package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 3
		m.textarea.SetWidth(msg.Width - 10)
		m.refreshViewport()

	case spinner.TickMsg:
		if m.loading {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	case catsMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.cats, m.totalItems, m.totalPages = msg.cats, msg.totalItems, msg.totalPages
			m.err = nil
			cmds = append(cmds, m.resolvePendingCommand()...)
		}
		m.refreshViewport()

	case categoriesMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.categories = msg.categories
			m.err = nil
			cmds = append(cmds, m.resolvePendingCommand()...)
		}
		m.refreshViewport()

	case tea.KeyMsg:
		switch m.state {
		case viewMenu:
			cmds = append(cmds, m.updateMenu(msg)...)
		case viewCats, viewCategories:
			cmds = append(cmds, m.updateList(msg)...)
		case viewChat:
			cmds = append(cmds, m.updateChat(msg)...)
		}
	}

	if m.state == viewCats || m.state == viewCategories {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
