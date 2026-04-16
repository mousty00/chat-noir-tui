package main

import (
	"fmt"
	"strings"
)

func renderWelcome() string {
	return "\n  " + catASCII + "\n\n  " +
		titleStyle.Render("Chat Noir — How can I help today?") + "\n\n  " +
		"Available commands:\n" +
		"  /search <query> — search cats by name or color\n" +
		"  /count          — show total counts\n" +
		"  /help           — show this help"

}

func renderMenu(cursor int) string {
	var b strings.Builder
	b.WriteString("\n  " + catASCII + "\n")
	b.WriteString("  " + titleStyle.Render("Chat Noir") + "\n\n")

	for i, entry := range menuEntries {
		if i == cursor {
			b.WriteString(menuSelectedStyle.Render("▶  "+entry.label) + "\n")
		} else {
			b.WriteString(menuItemStyle.Render("   "+entry.label) + "\n")
		}
		b.WriteString(menuDescStyle.Render(entry.desc) + "\n\n")
	}

	b.WriteString("\n  " + dimStyle.Render("↑/↓ navigate  •  Enter select  •  q quit"))
	return b.String()
}

func renderCatsView(m *model) string {
	if s, early := guardView(m.loading, m.spinner.View(), "Fetching cats…", m.err, len(m.cats) == 0, "No cats found."); early {
		return s
	}

	const nameW, colorW, catW = 20, 14, 18
	srcW := max(10, m.width-nameW-colorW-catW-10)

	var b strings.Builder
	fmt.Fprintf(&b, "\n  %s  %s\n\n",
		titleStyle.Render("Cats"),
		dimStyle.Render(fmt.Sprintf("%d total  •  page 1/%d", m.totalItems, m.totalPages)),
	)
	b.WriteString(tableHeaderStyle.Render(
		fmt.Sprintf("  %-*s  %-*s  %-*s  %-*s", nameW, "Name", colorW, "Color", catW, "Category", srcW, "Source"),
	) + "\n")

	for _, c := range m.cats {
		fmt.Fprintf(&b, "  %s  %s  %s  %s\n",
			cellStyle.Render(pad(c.Name, nameW)),
			cellMutedStyle.Render(pad(c.Color, colorW)),
			cellMutedStyle.Render(pad(orDash(c.Category.Name), catW)),
			cellMutedStyle.Render(pad(orDash(c.SourceName), srcW)),
		)
	}

	b.WriteString("\n  " + dimStyle.Render("↑/↓ scroll  •  Esc back"))
	return b.String()
}

func renderCategoriesView(m *model) string {
	if s, early := guardView(m.loading, m.spinner.View(), "Fetching categories…", m.err, len(m.categories) == 0, "No categories found."); early {
		return s
	}

	const nameW, hintW = 24, 16

	var b strings.Builder
	fmt.Fprintf(&b, "\n  %s  %s\n\n",
		titleStyle.Render("Cat Categories"),
		dimStyle.Render(fmt.Sprintf("%d total", len(m.categories))),
	)
	b.WriteString(tableHeaderStyle.Render(
		fmt.Sprintf("  %-*s  %-*s  %s", nameW, "Name", hintW, "Media Type", "ID"),
	) + "\n")

	for _, c := range m.categories {
		fmt.Fprintf(&b, "  %s  %s  %s\n",
			cellStyle.Render(pad(c.Name, nameW)),
			cellMutedStyle.Render(pad(c.MediaTypeHint, hintW)),
			cellMutedStyle.Render(shortID(c.ID)),
		)
	}

	b.WriteString("\n  " + dimStyle.Render("↑/↓ scroll  •  Esc back"))
	return b.String()
}

func renderChatView(m *model) string {
	if len(m.messages) == 0 {
		return renderWelcome()
	}

	var b strings.Builder
	for _, msg := range m.messages {
		role := aiStyle.Render("Assistant")
		if msg.role == "You" {
			role = userStyle.Render("User")
		}
		fmt.Fprintf(&b, " %s\n %s\n\n", role, contentStyle.Render(msg.content))
	}

	if m.loading {
		fmt.Fprintf(&b, " %s\n %s\n",
			aiStyle.Render("Assistant"),
			m.spinner.View()+" "+dimStyle.Render("Fetching data…"),
		)
	}

	return b.String()
}

func guardView(loading bool, spinnerView, loadMsg string, err error, empty bool, emptyMsg string) (string, bool) {
	switch {
	case loading:
		return "\n  " + spinnerView + "  " + loadMsg, true
	case err != nil:
		return "\n  " + errorStyle.Render("Error: "+err.Error()), true
	case empty:
		return "\n  " + dimStyle.Render(emptyMsg), true
	}
	return "", false
}

func pad(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n-1]) + "…"
	}
	return s + strings.Repeat(" ", n-len(runes))
}

func orDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func shortID(id string) string {
	if len(id) > 8 {
		return id[:8] + "…"
	}
	return id
}
