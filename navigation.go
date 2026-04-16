package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) updateMenu(msg tea.KeyMsg) []tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q":
		return []tea.Cmd{tea.Quit}
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
			m.refreshViewport()
		}
	case "down", "j":
		if m.cursor < len(menuEntries)-1 {
			m.cursor++
			m.refreshViewport()
		}
	case "enter", " ":
		return m.selectMenuItem()
	}
	return nil
}

func (m *model) selectMenuItem() []tea.Cmd {
	selected := menuEntries[m.cursor]
	m.state = selected.view
	m.err = nil

	switch selected.view {
	case viewCats:
		m.cats = nil
		return m.startFetch(fetchCatsCmd())
	case viewCategories:
		m.categories = nil
		return m.startFetch(fetchCategoriesCmd())
	case viewChat:
		m.refreshViewport()
		return []tea.Cmd{m.textarea.Focus()}
	}
	return nil
}

func (m *model) updateList(msg tea.KeyMsg) []tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q":
		return []tea.Cmd{tea.Quit}
	case "esc", "backspace":
		m.state = viewMenu
		m.refreshViewport()
	}
	return nil
}

func (m *model) updateChat(msg tea.KeyMsg) []tea.Cmd {
	switch msg.Type {
	case tea.KeyCtrlC:
		return []tea.Cmd{tea.Quit}

	case tea.KeyEsc:
		m.state = viewMenu
		m.textarea.Blur()
		m.refreshViewport()
		return nil

	case tea.KeyEnter:
		input := m.textarea.Value()
		if input == "" {
			break
		}
		m.reply("You", input)
		m.textarea.Reset()
		m.textarea.Blur()

		if strings.HasPrefix(input, "/") {
			return m.handleCommand(input)
		}
		m.reply("AI", "Only commands accepted. Type /help to see available commands.")
		return m.refocus()
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return []tea.Cmd{cmd}
}

func (m *model) refreshViewport() {
	switch m.state {
	case viewMenu:
		m.viewport.SetContent(renderMenu(m.cursor))
	case viewCats:
		m.viewport.SetContent(renderCatsView(m))
	case viewCategories:
		m.viewport.SetContent(renderCategoriesView(m))
	case viewChat:
		m.viewport.SetContent(renderChatView(m))
		if len(m.messages) > 0 {
			m.viewport.GotoBottom()
		}
	}
}

func (m *model) reply(role, content string) {
	m.messages = append(m.messages, chatMessage{role: role, content: content})
}

func (m *model) startFetch(cmds ...tea.Cmd) []tea.Cmd {
	m.loading = true
	m.refreshViewport()
	return append(cmds, m.spinner.Tick)
}

func (m *model) resolvePendingCommand() []tea.Cmd {
	if m.pendingCommand == "" || m.state != viewChat {
		return nil
	}
	return m.handleCommand(m.pendingCommand)
}
