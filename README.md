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

### Pre-built binaries (recommended)

Download the latest release for your platform from the [Releases page](../../releases/latest).

#### macOS (Apple Silicon)

```bash
tar -xzf chatnoir_Darwin_arm64.tar.gz
sudo mv chatnoir /usr/local/bin/
```

#### macOS (Intel)

```bash
tar -xzf chatnoir_Darwin_x86_64.tar.gz
sudo mv chatnoir /usr/local/bin/
```

#### Linux (x86_64)

```bash
tar -xzf chatnoir_Linux_x86_64.tar.gz
sudo mv chatnoir /usr/local/bin/
```

#### Linux (ARM64)

```bash
tar -xzf chatnoir_Linux_arm64.tar.gz
sudo mv chatnoir /usr/local/bin/
```

#### Windows (x86_64)

1. Extract `chatnoir_Windows_x86_64.zip`
2. Move `chatnoir.exe` to a folder in your `PATH`, e.g. `C:\Windows\System32\`

### From source

Requires Go 1.21+.

```bash
go install github.com/mousty00/chat-noir-tui@latest
```

## Usage

```bash
chatnoir
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
