#!/bin/bash

# colors
BOLD='\033[1m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[0;33m'
RESET='\033[0m'

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; fi
if [ "$ARCH" = "arm64" ]; then ARCH="arm64"; fi

VERSION=$(curl -fsSL https://api.github.com/repos/gurliv21/cmdock/releases/latest | grep '"tag_name"' | cut -d'"' -f4 | tr -d 'v')
ARCHIVE="cmdock_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/gurliv21/cmdock/releases/download/v${VERSION}/${ARCHIVE}"

echo ""
echo -e "${BOLD}cmdock installer${RESET}"
echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "  version  : ${CYAN}v${VERSION}${RESET}"
echo -e "  platform : ${CYAN}${OS}/${ARCH}${RESET}"
echo -e "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

echo -e "  ${YELLOW}→${RESET} Downloading..."
curl -fsSL "$DOWNLOAD_URL" -o /tmp/cmdock.tar.gz

echo -e "  ${YELLOW}→${RESET} Extracting..."
tar -xzf /tmp/cmdock.tar.gz -C /tmp
chmod +x /tmp/cmdock

echo -e "  ${YELLOW}→${RESET} Installing to /usr/local/bin..."
sudo mv /tmp/cmdock /usr/local/bin/cmdock
rm /tmp/cmdock.tar.gz

echo ""
echo -e "  ${GREEN}✓${RESET} ${BOLD}cmdock installed successfully!${RESET}"
echo ""
echo -e "  Next steps:"
echo -e "  ${CYAN}1.${RESET} cmdock init"
echo -e "  ${CYAN}2.${RESET} source ~/.zshrc"
echo ""