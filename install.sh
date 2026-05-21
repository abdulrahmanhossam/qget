#!/bin/bash

set -e

REPO="abdulrahmanhossam/qget"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="qget"

echo "==> Fetching latest release for ${REPO}..."

TAG=$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | cut -d'"' -f4)

if [ -z "$TAG" ]; then
    echo "Failed to fetch latest release tag."
    exit 1
fi

echo "==> Latest version: ${TAG}"

TEMP_DIR=$(mktemp -d)
trap "rm -rf ${TEMP_DIR}" EXIT

echo "==> Downloading qget-linux..."

curl -sL "https://github.com/${REPO}/releases/download/${TAG}/qget-linux" -o "${TEMP_DIR}/${BINARY_NAME}"

echo "==> Making binary executable..."
chmod +x "${TEMP_DIR}/${BINARY_NAME}"

echo "==> Moving to ${INSTALL_DIR}/${BINARY_NAME}..."
sudo mv "${TEMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"

echo "==> Installation complete! Run 'qget --help' to get started."