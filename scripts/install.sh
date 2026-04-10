#!/bin/bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# normalize arch
if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; fi
if [ "$ARCH" = "arm64" ]; then ARCH="arm64"; fi

VERSION="1.0.6"
BINARY_URL="https://github.com/gurlivbajwa/cmdock/releases/download/v${VERSION}/cmdock_${OS}_${ARCH}"

echo "Installing cmdock..."

curl -fsSL "$BINARY_URL" -o /tmp/cmdock
chmod +x /tmp/cmdock
sudo mv /tmp/cmdock /usr/local/bin/cmdock

echo " cmdock installed to /usr/local/bin/cmdock"
echo " Run: cmdock init"
echo " Then: source ~/.zshrc"