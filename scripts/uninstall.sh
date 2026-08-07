#!/usr/bin/env sh
set -e

BINARY="cmdock"
INSTALL_DIRS="/usr/local/bin $HOME/.local/bin"

# --- colors ---
if [ -t 1 ]; then
    BOLD="$(printf '\033[1m')"
    GREEN="$(printf '\033[32m')"
    CYAN="$(printf '\033[36m')"
    RED="$(printf '\033[31m')"
    DIM="$(printf '\033[2m')"
    RESET="$(printf '\033[0m')"
else
    BOLD="" GREEN="" CYAN="" RED="" DIM="" RESET=""
fi

info()    { printf "%s%s%s\n" "$CYAN" "$1" "$RESET"; }
success() { printf "%s✓ %s%s\n" "$GREEN" "$1" "$RESET"; }
warn()    { printf "%s! %s%s\n" "$DIM" "$1" "$RESET"; }
error()   { printf "%s✗ %s%s\n" "$RED" "$1" "$RESET" >&2; }

printf "\n%s%s cmdock uninstaller %s\n\n" "$BOLD" "▸" "$RESET"

printf "%sThis will remove the cmdock binary, database, and shell hooks. Continue? [y/N]%s " "$CYAN" "$RESET"
read -r ANSWER
case "$ANSWER" in
    [yY]*) ;;
    *)
        printf "Cancelled.\n"
        exit 0
        ;;
esac
printf "\n"

# --- 1. remove binary ---
REMOVED_BIN=0
for dir in $INSTALL_DIRS; do
    if [ -f "$dir/$BINARY" ]; then
        if [ -w "$dir" ]; then
            rm -f "$dir/$BINARY"
        else
            sudo rm -f "$dir/$BINARY"
        fi
        success "Removed binary: $dir/$BINARY"
        REMOVED_BIN=1
    fi
done
[ "$REMOVED_BIN" -eq 0 ] && warn "No binary found in known install locations"

# --- 2. remove database ---
DB_PATH="$HOME/.cmdock.db"
if [ -f "$DB_PATH" ]; then
    rm -f "$DB_PATH"
    success "Database deleted"
else
    warn "No database found"
fi

# --- 3. strip shell hooks ---
strip_block() {
    file="$1"
    [ -f "$file" ] || return
    if grep -q "# >>> cmdock start >>>" "$file" 2>/dev/null; then
        sed -i.bak '/# >>> cmdock start >>>/,/# <<< cmdock end <<</d' "$file"
        rm -f "$file.bak"
        success "Shell integration removed ($file)"
    fi
}

strip_block "$HOME/.bashrc"
strip_block "$HOME/.zshrc"
strip_block "$HOME/.config/fish/config.fish"

printf "\n%s✓ cmdock fully uninstalled%s\n" "$GREEN$BOLD" "$RESET"
printf "%sRestart your terminal, or run: source ~/.zshrc (or ~/.bashrc)%s\n\n" "$DIM" "$RESET"