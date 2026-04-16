# chat-noir-tui

Terminal UI for [Chat Noir](https://chatnoir.fun) — browse cats and categories, search by name or color, and chat via commands. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

```
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
```

## Install

Requires Go 1.21+.

```bash
go install github.com/mousty00/chat-noir-tui@latest
```

## Usage

```bash
chat-noir-tui
```

## Navigation

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Enter` | Select |
| `Esc` / `Backspace` | Back to menu |
| `q` / `Ctrl+C` | Quit |

## Views

- **Cats** — browse all cats fetched from the API
- **Categories** — browse cat categories
- **Chat** — command-based interface

## Chat Commands

| Command | Description |
|---------|-------------|
| `/search <query>` | Search cats by name or color |
| `/count` | Show total cat and category counts |
| `/help` | List available commands |

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) — UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — styling
