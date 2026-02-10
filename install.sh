#!/bin/bash

# Thrustlab GTM Skills + MCP Servers — Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/kiwiidb/thrustlab-skills-library/main/install.sh | bash

set -e

REPO="kiwiidb/thrustlab-skills-library"
CLAUDE_SKILLS_DIR="$HOME/.claude/skills"
CLAUDE_BIN_DIR="$HOME/.claude/bin"
CLAUDE_DESKTOP_CONFIG="$HOME/Library/Application Support/Claude/claude_desktop_config.json"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BOLD='\033[1m'
NC='\033[0m'

info()  { echo -e "${GREEN}✓${NC} $1"; }
warn()  { echo -e "${YELLOW}!${NC} $1"; }
error() { echo -e "${RED}✗${NC} $1"; exit 1; }

echo ""
echo -e "${BOLD}Thrustlab GTM — Skills & MCP Installer${NC}"
echo "────────────────────────────────────────"
echo ""

# ── Fetch latest release ─────────────────────────────────────────────

echo "Fetching latest release..."
RELEASE_JSON=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest") || error "Failed to fetch release info. Is the repo public?"

VERSION=$(echo "$RELEASE_JSON" | python3 -c "import sys,json; print(json.load(sys.stdin)['tag_name'])")
TARBALL_URL=$(echo "$RELEASE_JSON" | python3 -c "
import sys, json
assets = json.load(sys.stdin)['assets']
for a in assets:
    if a['name'].endswith('.tar.gz'):
        print(a['browser_download_url'])
        break
")

if [ -z "$TARBALL_URL" ]; then
    error "No tarball found in release $VERSION"
fi

info "Latest release: $VERSION"

# ── Download and extract ─────────────────────────────────────────────

TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

echo "Downloading $VERSION..."
curl -fsSL -o "$TEMP_DIR/thrustlab.tar.gz" "$TARBALL_URL" || error "Download failed"
tar -xzf "$TEMP_DIR/thrustlab.tar.gz" -C "$TEMP_DIR" || error "Extract failed"

info "Downloaded and extracted"

# ── Install skills ───────────────────────────────────────────────────

mkdir -p "$CLAUDE_SKILLS_DIR"
cp -r "$TEMP_DIR/skills/"* "$CLAUDE_SKILLS_DIR/"
SKILL_COUNT=$(find "$CLAUDE_SKILLS_DIR" -maxdepth 1 -type d | tail -n +2 | wc -l | tr -d ' ')
info "Installed $SKILL_COUNT skills → $CLAUDE_SKILLS_DIR"

# ── Install MCP binaries ────────────────────────────────────────────

mkdir -p "$CLAUDE_BIN_DIR"
cp "$TEMP_DIR/bin/"* "$CLAUDE_BIN_DIR/"
chmod +x "$CLAUDE_BIN_DIR/"*
info "Installed MCP binaries → $CLAUDE_BIN_DIR"

# ── Register MCPs ────────────────────────────────────────────────────

MCP_SERVERS=(
    "slack:slack-mcp"
    "clay:clay-mcp-server"
    "namecheap:namecheap-mcp"
    "premiuminboxes:premiuminboxes-mcp"
)

if command -v claude &> /dev/null; then
    echo ""
    echo "Registering MCP servers with Claude Code..."
    for entry in "${MCP_SERVERS[@]}"; do
        IFS=':' read -r name binary <<< "$entry"
        if claude mcp get "$name" &> /dev/null 2>&1; then
            warn "$name — already registered, skipping"
        else
            claude mcp add -s user "$name" -- "$CLAUDE_BIN_DIR/$binary" 2>/dev/null && \
                info "$name — registered" || \
                warn "$name — could not register (add manually later)"
        fi
    done
fi

# ── Patch Claude Desktop config (fallback) ───────────────────────────

if [ -d "$(dirname "$CLAUDE_DESKTOP_CONFIG")" ]; then
    echo ""
    echo "Updating Claude Desktop configuration..."
    python3 -c "
import json, os

config_path = os.path.expanduser('~/Library/Application Support/Claude/claude_desktop_config.json')
config = {}
if os.path.exists(config_path):
    with open(config_path) as f:
        config = json.load(f)

bin_dir = os.path.expanduser('~/.claude/bin')
servers = config.setdefault('mcpServers', {})

mcp_defs = {
    'slack': {'command': f'{bin_dir}/slack-mcp', 'env': {'SLACK_BOT_TOKEN': servers.get('slack', {}).get('env', {}).get('SLACK_BOT_TOKEN', '')}},
    'clay': {'command': f'{bin_dir}/clay-mcp-server', 'env': {'CLAY_WORKSPACE_ID': servers.get('clay', {}).get('env', {}).get('CLAY_WORKSPACE_ID', ''), 'CLAY_SESSION_COOKIE': servers.get('clay', {}).get('env', {}).get('CLAY_SESSION_COOKIE', '')}},
    'namecheap': {'command': f'{bin_dir}/namecheap-mcp', 'env': {'NAMECHEAP_API_USER': servers.get('namecheap', {}).get('env', {}).get('NAMECHEAP_API_USER', ''), 'NAMECHEAP_API_KEY': servers.get('namecheap', {}).get('env', {}).get('NAMECHEAP_API_KEY', ''), 'NAMECHEAP_USERNAME': servers.get('namecheap', {}).get('env', {}).get('NAMECHEAP_USERNAME', ''), 'NAMECHEAP_CLIENT_IP': servers.get('namecheap', {}).get('env', {}).get('NAMECHEAP_CLIENT_IP', '')}},
    'premiuminboxes': {'command': f'{bin_dir}/premiuminboxes-mcp', 'env': {'PREMIUMINBOXES_API_TOKEN': servers.get('premiuminboxes', {}).get('env', {}).get('PREMIUMINBOXES_API_TOKEN', '')}},
}

for name, defn in mcp_defs.items():
    if name not in servers:
        servers[name] = defn

with open(config_path, 'w') as f:
    json.dump(config, f, indent=2)
" && info "Claude Desktop config updated" || warn "Could not update Claude Desktop config"
fi

# ── Summary ──────────────────────────────────────────────────────────

echo ""
echo -e "${BOLD}Installation complete!${NC} ($VERSION)"
echo "────────────────────────────────────────"
echo ""
echo "Skills installed:"
for skill in "$CLAUDE_SKILLS_DIR"/*/; do
    [ -d "$skill" ] && echo "  /${BOLD}$(basename "$skill")${NC}"
done
echo ""
echo "MCP servers installed:"
for entry in "${MCP_SERVERS[@]}"; do
    IFS=':' read -r name binary <<< "$entry"
    echo "  $name → $CLAUDE_BIN_DIR/$binary"
done
echo ""
echo -e "${YELLOW}Next step: Configure API credentials${NC}"
echo ""
echo "  Slack:"
echo "    claude mcp add -e SLACK_BOT_TOKEN=\"xoxb-...\" -s user slack -- $CLAUDE_BIN_DIR/slack-mcp"
echo ""
echo "  Clay:"
echo "    claude mcp add -e CLAY_WORKSPACE_ID=\"...\" -e CLAY_SESSION_COOKIE=\"...\" -s user clay -- $CLAUDE_BIN_DIR/clay-mcp-server"
echo ""
echo "  Namecheap:"
echo "    claude mcp add -e NAMECHEAP_API_USER=\"...\" -e NAMECHEAP_API_KEY=\"...\" -e NAMECHEAP_USERNAME=\"...\" -e NAMECHEAP_CLIENT_IP=\"...\" -s user namecheap -- $CLAUDE_BIN_DIR/namecheap-mcp"
echo ""
echo "  Premium Inboxes:"
echo "    claude mcp add -e PREMIUMINBOXES_API_TOKEN=\"...\" -s user premiuminboxes -- $CLAUDE_BIN_DIR/premiuminboxes-mcp"
echo ""
echo "Restart Claude Desktop after configuring credentials."
echo ""
