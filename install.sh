#!/bin/bash

# Thrustlab GTM Claude Skills - Installation Script
# This script installs Thrustlab skills to your Claude skills directory

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect Claude skills directory
CLAUDE_SKILLS_DIR="$HOME/.claude/skills"

echo "╔════════════════════════════════════════════════════════╗"
echo "║   Thrustlab GTM Claude Skills - Installer             ║"
echo "╚════════════════════════════════════════════════════════╝"
echo ""

# Check if Claude directory exists
if [ ! -d "$HOME/.claude" ]; then
    echo -e "${YELLOW}Warning: ~/.claude directory not found.${NC}"
    echo "This might mean Claude Desktop/Code is not installed."
    echo ""
    read -p "Do you want to create the directory anyway? (y/n) " -n 1 -r
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${RED}Installation cancelled.${NC}"
        exit 1
    fi
fi

# Create skills directory if it doesn't exist
mkdir -p "$CLAUDE_SKILLS_DIR"
echo -e "${GREEN}✓${NC} Skills directory ready: $CLAUDE_SKILLS_DIR"

# Determine installation method
if [ -d ".git" ] && [ -d "skills" ]; then
    # Running from cloned repository
    echo ""
    echo "Installing from local repository..."
    cp -r skills/* "$CLAUDE_SKILLS_DIR/"
    echo -e "${GREEN}✓${NC} Skills copied to $CLAUDE_SKILLS_DIR"
else
    # Need to download from GitHub
    echo ""
    echo "Downloading latest skills from GitHub..."

    TEMP_DIR=$(mktemp -d)

    # Download and extract
    if command -v curl &> /dev/null; then
        curl -fsSL https://github.com/your-org/thrustlab/archive/main.tar.gz | tar -xz -C "$TEMP_DIR"
    elif command -v wget &> /dev/null; then
        wget -qO- https://github.com/your-org/thrustlab/archive/main.tar.gz | tar -xz -C "$TEMP_DIR"
    else
        echo -e "${RED}Error: curl or wget is required but not installed.${NC}"
        exit 1
    fi

    # Copy skills
    cp -r "$TEMP_DIR"/thrustlab-main/skills/* "$CLAUDE_SKILLS_DIR/"

    # Cleanup
    rm -rf "$TEMP_DIR"

    echo -e "${GREEN}✓${NC} Skills downloaded and installed"
fi

echo ""
echo "╔════════════════════════════════════════════════════════╗"
echo "║   Installation Complete!                               ║"
echo "╚════════════════════════════════════════════════════════╝"
echo ""
echo "Installed skills:"
echo "  • notion-client-onboarding"
echo ""
echo -e "${YELLOW}Important: Some skills require MCP servers${NC}"
echo ""
echo "For Notion Client Onboarding skill, you need:"
echo "  1. Install Notion MCP Server"
echo "  2. Configure in Claude Desktop settings"
echo "  3. See README.md for detailed instructions"
echo ""
echo "Documentation: https://github.com/your-org/thrustlab"
echo ""
echo -e "${GREEN}Happy GTM planning! 🚀${NC}"
