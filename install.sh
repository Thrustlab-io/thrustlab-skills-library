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

# Install Slack MCP globally
echo ""
echo "Installing Slack MCP..."

CLAUDE_BIN_DIR="$HOME/.claude/bin"
mkdir -p "$CLAUDE_BIN_DIR"

if [ -d ".git" ] && [ -f "slack-mcp/slack-mcp" ]; then
    # Copy binary to global location
    cp slack-mcp/slack-mcp "$CLAUDE_BIN_DIR/"
    chmod +x "$CLAUDE_BIN_DIR/slack-mcp"
    echo -e "${GREEN}✓${NC} Slack MCP binary installed to $CLAUDE_BIN_DIR"

    # Register with Claude (requires SLACK_BOT_TOKEN to be set)
    if command -v claude &> /dev/null; then
        if claude mcp get slack &> /dev/null; then
            echo -e "${YELLOW}!${NC} Slack MCP already registered with Claude"
        else
            if [ -n "$SLACK_BOT_TOKEN" ]; then
                claude mcp add -e SLACK_BOT_TOKEN="$SLACK_BOT_TOKEN" -s user slack -- "$CLAUDE_BIN_DIR/slack-mcp"
                echo -e "${GREEN}✓${NC} Slack MCP registered with Claude"
            else
                echo -e "${YELLOW}!${NC} Set SLACK_BOT_TOKEN and run: claude mcp add -e SLACK_BOT_TOKEN=\"\$SLACK_BOT_TOKEN\" -s user slack -- $CLAUDE_BIN_DIR/slack-mcp"
            fi
        fi
    fi
else
    echo -e "${YELLOW}!${NC} Slack MCP binary not found (run from cloned repo to install)"
fi

# Install Premium Inboxes MCP
echo ""
echo "Installing Premium Inboxes MCP..."

if [ -d ".git" ] && [ -d "premiuminboxes-mcp" ]; then
    # Build if needed
    if [ ! -f "premiuminboxes-mcp/premiuminboxes-mcp" ]; then
        echo "Building Premium Inboxes MCP..."
        if command -v go &> /dev/null; then
            (cd premiuminboxes-mcp && go build -o premiuminboxes-mcp)
            echo -e "${GREEN}✓${NC} Premium Inboxes MCP built successfully"
        else
            echo -e "${RED}✗${NC} Go is required to build Premium Inboxes MCP"
            echo "Please install Go from https://golang.org/dl/"
        fi
    fi

    if [ -f "premiuminboxes-mcp/premiuminboxes-mcp" ]; then
        # Copy binary to global location
        cp premiuminboxes-mcp/premiuminboxes-mcp "$CLAUDE_BIN_DIR/"
        chmod +x "$CLAUDE_BIN_DIR/premiuminboxes-mcp"
        echo -e "${GREEN}✓${NC} Premium Inboxes MCP binary installed to $CLAUDE_BIN_DIR"

        # Register with Claude (requires PREMIUMINBOXES_API_TOKEN to be set)
        if command -v claude &> /dev/null; then
            if claude mcp get premiuminboxes &> /dev/null; then
                echo -e "${YELLOW}!${NC} Premium Inboxes MCP already registered with Claude"
            else
                if [ -n "$PREMIUMINBOXES_API_TOKEN" ]; then
                    claude mcp add -e PREMIUMINBOXES_API_TOKEN="$PREMIUMINBOXES_API_TOKEN" -s user premiuminboxes -- "$CLAUDE_BIN_DIR/premiuminboxes-mcp"
                    echo -e "${GREEN}✓${NC} Premium Inboxes MCP registered with Claude"
                else
                    echo -e "${YELLOW}!${NC} Set PREMIUMINBOXES_API_TOKEN and run: claude mcp add -e PREMIUMINBOXES_API_TOKEN=\"\$PREMIUMINBOXES_API_TOKEN\" -s user premiuminboxes -- $CLAUDE_BIN_DIR/premiuminboxes-mcp"
                fi
            fi
        fi
    fi
else
    echo -e "${YELLOW}!${NC} Premium Inboxes MCP not found (run from cloned repo to install)"
fi

# Install Clay MCP
echo ""
echo "Installing Clay MCP..."

if [ -d ".git" ] && [ -d "clay-mcp-server" ]; then
    # Build if needed
    if [ ! -f "clay-mcp-server/clay-mcp-server" ]; then
        echo "Building Clay MCP..."
        if command -v go &> /dev/null; then
            (cd clay-mcp-server && go build -o clay-mcp-server)
            echo -e "${GREEN}✓${NC} Clay MCP built successfully"
        else
            echo -e "${RED}✗${NC} Go is required to build Clay MCP"
            echo "Please install Go from https://golang.org/dl/"
        fi
    fi

    if [ -f "clay-mcp-server/clay-mcp-server" ]; then
        # Copy binary to global location
        cp clay-mcp-server/clay-mcp-server "$CLAUDE_BIN_DIR/"
        chmod +x "$CLAUDE_BIN_DIR/clay-mcp-server"
        echo -e "${GREEN}✓${NC} Clay MCP binary installed to $CLAUDE_BIN_DIR"

        # Register with Claude (requires CLAY_WORKSPACE_ID and CLAY_SESSION_COOKIE)
        if command -v claude &> /dev/null; then
            if claude mcp get clay &> /dev/null; then
                echo -e "${YELLOW}!${NC} Clay MCP already registered with Claude"
            else
                if [ -n "$CLAY_WORKSPACE_ID" ] && [ -n "$CLAY_SESSION_COOKIE" ]; then
                    claude mcp add -e CLAY_WORKSPACE_ID="$CLAY_WORKSPACE_ID" -e CLAY_SESSION_COOKIE="$CLAY_SESSION_COOKIE" -s user clay -- "$CLAUDE_BIN_DIR/clay-mcp-server"
                    echo -e "${GREEN}✓${NC} Clay MCP registered with Claude"
                else
                    echo -e "${YELLOW}!${NC} To register Clay MCP, set credentials and run:"
                    echo "    claude mcp add -e CLAY_WORKSPACE_ID=\"your-workspace-id\" -e CLAY_SESSION_COOKIE=\"your-session-cookie\" -s user clay -- $CLAUDE_BIN_DIR/clay-mcp-server"
                    echo ""
                    echo "  Or configure after installation using Claude commands:"
                    echo "    Use 'set_workspace_id' and 'set_session_cookie' tools in Claude"
                fi
            fi
        fi
    fi
else
    echo -e "${YELLOW}!${NC} Clay MCP not found (run from cloned repo to install)"
fi

echo ""
echo "╔════════════════════════════════════════════════════════╗"
echo "║   Installation Complete!                               ║"
echo "╚════════════════════════════════════════════════════════╝"
echo ""
echo "Installed skills:"
for skill in "$CLAUDE_SKILLS_DIR"/*; do
    if [ -d "$skill" ]; then
        skill_name=$(basename "$skill")
        echo "  • $skill_name"
    fi
done
echo ""
echo -e "${YELLOW}Important: Setup requirements${NC}"
echo ""
echo "MCP Setup:"
echo "  • Slack MCP: Run with SLACK_BOT_TOKEN set, or manually add:"
echo "    claude mcp add -e SLACK_BOT_TOKEN=\"your-token\" -s user slack -- ~/.claude/bin/slack-mcp"
echo ""
echo "  • Premium Inboxes MCP: Run with PREMIUMINBOXES_API_TOKEN set, or manually add:"
echo "    claude mcp add -e PREMIUMINBOXES_API_TOKEN=\"your-token\" -s user premiuminboxes -- ~/.claude/bin/premiuminboxes-mcp"
echo ""
echo "  • Clay MCP: Run with CLAY_WORKSPACE_ID and CLAY_SESSION_COOKIE set, or manually add:"
echo "    claude mcp add -e CLAY_WORKSPACE_ID=\"your-id\" -e CLAY_SESSION_COOKIE=\"your-cookie\" -s user clay -- ~/.claude/bin/clay-mcp-server"
echo "    Or configure after installation using: set_workspace_id and set_session_cookie tools"
echo ""
echo "Some skills may require additional setup:"
echo "  • External service accounts (Clay, Notion, Slack, Premium Inboxes)"
echo ""
echo "See README.md and individual SKILL.md files for detailed requirements"
echo ""
echo "Documentation: https://github.com/kwinten/thrustlab"
echo ""
echo -e "${GREEN}Happy GTM planning! 🚀${NC}"
