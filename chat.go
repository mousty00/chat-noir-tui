package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) handleCommand(input string) []tea.Cmd {
	m.pendingCommand = ""

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	name := strings.ToLower(parts[0])
	args := parts[1:]

	switch name {
	case "/help":
		m.reply("AI",
			"Available commands:\n"+
				"  /search <query> — search cats by name or color\n"+
				"  /count          — show total counts\n"+
				"  /help           — show this help",
		)

	case "/count":
		if m.cats == nil || m.categories == nil {
			m.reply("AI", "Fetching data, one moment…")
			m.pendingCommand = input
			return m.startFetch(fetchCatsCmd(), fetchCategoriesCmd())
		}
		m.reply("AI", fmt.Sprintf("Found %d cats across %d categories.", m.totalItems, len(m.categories)))

	case "/search":
		if len(args) == 0 {
			m.reply("AI", "Usage: /search <query>  —  example: /search black")
			break
		}
		query := strings.Join(args, " ")
		if m.cats == nil {
			m.reply("AI", "Fetching cats before searching…")
			m.pendingCommand = input
			return m.startFetch(fetchCatsCmd())
		}
		return m.performSearch(query)

	default:
		m.reply("AI", fmt.Sprintf("Unknown command %q. Type /help for assistance.", name))
	}

	return m.refocus()
}

func (m *model) performSearch(query string) []tea.Cmd {
	q := strings.ToLower(query)

	var matches []Cat
	for _, c := range m.cats {
		nameMatch := strings.Contains(strings.ToLower(c.Name), q)
		colorMatch := strings.Contains(strings.ToLower(c.Color), q)
		if nameMatch || colorMatch {
			matches = append(matches, c)
		}
	}

	if len(matches) == 0 {
		m.reply("AI", fmt.Sprintf("No cats found matching %q.", query))
		return m.refocus()
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Found %d cats matching %q:\n", len(matches), query)
	for i, c := range matches {
		if i == 5 {
			fmt.Fprintf(&b, "\n…and %d more. Open the Cats view for the full list.", len(matches)-5)
			break
		}
		fmt.Fprintf(&b, "\n• %s (%s)", c.Name, c.Color)
	}
	m.reply("AI", b.String())
	return m.refocus()
}

func (m *model) refocus() []tea.Cmd {
	m.refreshViewport()
	m.viewport.GotoBottom()
	return []tea.Cmd{m.textarea.Focus(), textarea.Blink}
}
