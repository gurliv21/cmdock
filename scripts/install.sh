#!/usr/bin/env sh
set -e

REPO="gurliv21/cmdock"
BINARY="cmdock"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

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
error()   { printf "%s✗ %s%s\n" "$RED" "$1" "$RESET" >&2; }

printf "\n%s%s cmdock installer %s\n\n" "$BOLD" "▸" "$RESET"

#os detection
OS="$(uname -s)"
case "$OS" in
    Linux)  OS="linux" ;;
    Darwin) OS="darwin" ;;
    *)
        error "Unsupported OS: $OS"
        printf "Download manually: %shttps://github.com/%s/releases%s\n" "$DIM" "$REPO" "$RESET"
        exit 1
        ;;
esac

ARCH="$(uname -m)"
case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *)
        error "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

info "Detected: ${OS}/${ARCH}"

printf "Fetching latest release...\n"
VERSION="$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')"

if [ -z "$VERSION" ]; then
    error "Failed to fetch latest version"
    exit 1
fi

success "Latest version: ${BOLD}${VERSION}${RESET}"


ARCHIVE="${BINARY}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/${VERSION}/${ARCHIVE}"


TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

printf "Downloading...\n"
curl -fsSL "$URL" -o "$TMP_DIR/$ARCHIVE" || {
    error "Download failed"
    exit 1
}
success "Downloaded"


tar -xzf "$TMP_DIR/$ARCHIVE" -C "$TMP_DIR"

if [ ! -f "$TMP_DIR/$BINARY" ]; then
    error "Binary not found in archive"
    exit 1
fi

chmod +x "$TMP_DIR/$BINARY"

if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/$BINARY"
else
    printf "%sNeed sudo to write to %s%s\n" "$DIM" "$INSTALL_DIR" "$RESET"
    sudo mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/$BINARY"
fi

printf "\n%s✓ cmdock installed successfully%s\n" "$GREEN$BOLD" "$RESET"
printf "%s→ %s%s\n\n" "$CYAN" "$BINARY --version" "$RESET"

if [ -t 0 ] && [ -t 1 ]; then
    printf "%sRun shell init now to enable command tracking? [Y/n]%s " "$CYAN" "$RESET"
    read -r ANSWER
    case "$ANSWER" in
        [nN]*)
            printf "\n%sSkipped. Run this later:%s\n" "$DIM" "$RESET"
            printf "  %s%s init%s\n\n" "$BOLD" "$BINARY" "$RESET"
            ;;
        *)
            printf "\n"
            "$INSTALL_DIR/$BINARY" init
            success "Shell initialized"
            printf "\n%sRestart your shell (or run 'source' on your profile) to start using cmdock.%s\n\n" "$DIM" "$RESET"
            ;;
    esac
else
    printf "%sNext step:%s\n" "$CYAN" "$RESET"
    printf "  %s%s init%s\n\n" "$BOLD" "$BINARY" "$RESET"
fi