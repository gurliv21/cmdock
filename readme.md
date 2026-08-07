# cmdock

`cmdock` is a lightweight command history manager for the terminal. It automatically records every shell command you execute and gives you an interactive TUI to browse, fuzzy-search, and inspect your history — including working directory, exit status, and timing for every command.

Built in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the interface and SQLite for persistent storage.

---

## Features

- Automatic shell command recording (bash, zsh, fish, PowerShell)
- Interactive terminal UI (TUI)
- Real-time fuzzy search
- Tracks working directory, exit code, and start/end time per command
- Persistent local SQLite database
- Fast, keyboard-driven navigation
- Single static binary, no runtime dependencies

---

## Installation

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/gurliv21/cmdock/main/scripts/install.sh | sh
```

This detects your OS/architecture, downloads the latest release, installs the binary to `/usr/local/bin`, and offers to run shell setup for you.

### Windows

```powershell
irm https://raw.githubusercontent.com/gurliv21/cmdock/main/scripts/install.ps1 | iex
```

Installs to `%LOCALAPPDATA%\cmdock` and adds it to your user `PATH` automatically.

### Manual download

Grab a prebuilt binary for your platform from the [Releases page](https://github.com/gurliv21/cmdock/releases), extract it, and move it somewhere on your `PATH`:

```bash
tar -xzf cmdock_<os>_<arch>.tar.gz
sudo mv cmdock /usr/local/bin
```

### Build from source

Requires Go 1.22+.

```bash
git clone https://github.com/gurliv21/cmdock.git
cd cmdock
go build -o cmdock .
```

or, if you have Go installed and just want the binary on your `PATH`:

```bash
go install github.com/gurliv21/cmdock@latest
```

---

## Shell integration

`cmdock` needs a small shell hook to capture commands as you run them. If you installed via `install.sh`/`install.ps1`, you were already offered this step. Otherwise, run:

```bash
cmdock init
```

This detects your current shell (bash, zsh, fish, or PowerShell) and appends the required hook to the right config file automatically. Restart your terminal, or reload your shell:

```bash
source ~/.zshrc     # zsh
source ~/.bashrc    # bash
```

---

## Usage

Launch the interface:

```bash
cmdock
```

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
