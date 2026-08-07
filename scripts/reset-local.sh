#!/usr/bin/env sh
set -e

if [ -t 1 ]; then
    GREEN="$(printf '\033[32m')"; CYAN="$(printf '\033[36m')"
    DIM="$(printf '\033[2m')"; RESET="$(printf '\033[0m')"
else
    GREEN="" CYAN="" DIM="" RESET=""
fi

log() { printf "%s✓ %s%s\n" "$GREEN" "$1" "$RESET"; }

printf "\n%sResetting local cmdock environment...%s\n\n" "$CYAN" "$RESET"

# --- remove binaries from every known location ---
for path in /usr/local/bin/cmdock "$(go env GOPATH 2>/dev/null)/bin/cmdock" "$HOME/.local/bin/cmdock"; do
    if [ -f "$path" ]; then
        if [ -w "$(dirname "$path")" ]; then
            rm -f "$path"
        else
            sudo rm -f "$path"
        fi
        log "Removed binary: $path"
    fi
done

# --- remove database ---
if [ -f "$HOME/.cmdock.db" ]; then
    rm -f "$HOME/.cmdock.db"
    log "Removed database"
fi

# --- strip hooks from all shells ---
for rc in "$HOME/.zshrc" "$HOME/.bashrc" "$HOME/.config/fish/config.fish"; do
    if [ -f "$rc" ] && grep -q "# >>> cmdock start >>>" "$rc" 2>/dev/null; then
        sed -i.bak '/# >>> cmdock start >>>/,/# <<< cmdock end <<</d' "$rc"
        rm -f "$rc.bak"
        log "Cleaned hook: $rc"
    fi
done

printf "\n%s✓ Fully reset. Ready for a fresh install.%s\n" "$GREEN" "$RESET"
printf "%sRun 'which -a cmdock' to confirm nothing is left.%s\n\n" "$DIM" "$RESET"