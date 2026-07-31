# cmdock

`cmdock` is a lightweight command history manager for the terminal. It automatically records every shell command you execute and provides an interactive terminal interface to browse, search, and inspect your history.

The project is built in Go using Bubble Tea for the terminal UI and SQLite for persistent storage.

---

## Features

* Automatic shell command recording
* Interactive terminal user interface (TUI)
* Real-time fuzzy search
* Stores the working directory for each command
* Records command exit status
* Records execution start and end times
* Persistent SQLite database
* Fast keyboard-driven navigation

---

## Installation

### Clone the repository

```bash
git clone https://github.com/<username>/cmdock.git
cd cmdock
```

### Build

```bash
go build -o cmdock
```

or install with Go:

```bash
go install
```

---

## Shell Integration

`cmdock` records commands by integrating with your shell. Add the provided shell hook to your shell configuration file.

### Zsh

Add the hook to your `~/.zshrc`.

```sh
# cmdock shell hook
# (Add the provided hook here)
```

Reload your shell:

```bash
source ~/.zshrc
```

### Bash

Add the hook to your `~/.bashrc`.

```sh
# cmdock shell hook
```

Reload your shell:

```bash
source ~/.bashrc
```

---

## Running

Launch the interface with:

```bash
cmdock
```

---

## Keyboard Shortcuts

| Key   | Action                   |
| ----- | ------------------------ |
| ↑ / ↓ | Navigate command history |
| Enter | Open command details     |
| /     | Start fuzzy search       |
| Esc   | Close search or details  |
| q     | Quit                     |

---

## Search

Press `/` to begin searching.

Search is performed using fuzzy matching and updates results as you type. Partial and abbreviated queries are supported.

Examples:

```
gst
→ git status

gcm
→ git commit -m

dcu
→ docker compose up
```

No additional key press is required—the displayed results update immediately on every keystroke.

---

## Stored Information

For every recorded command, `cmdock` stores:

* Command
* Working directory
* Exit code
* Start time
* End time

---

## Project Structure

```
cmdock/
├── cmd/
├── internal/
│   ├── store/
│   └── tui/
├── main.go
└── README.md
```

---

## Built With

* Go
* Bubble Tea
* Bubbles
* Lip Gloss
* SQLite
* sahilm/fuzzy

---

## Roadmap

* Interactive command history browser
* Real-time fuzzy search
* Clipboard support
* Command re-execution
* Favorites and pinned commands
* Command statistics
* Import existing shell history
* Configurable themes

---

## License

This project is licensed under the MIT License.
