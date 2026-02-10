#!/bin/bash

# Thrustlab GTM Skills + MCP Servers — Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/Thrustlab-io/thrustlab-skills-library/main/install.sh | bash

set -e

REPO="Thrustlab-io/thrustlab-skills-library"
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

# Read from /dev/tty so prompts work even when piped from curl
prompt() {
    local var="$1" label="$2"
    printf "  %s: " "$label" > /dev/tty
    read -r "$var" < /dev/tty
}

echo ""
echo -e "${BOLD}Thrustlab GTM — Skills & MCP Installer${NC}"
echo "────────────────────────────────────────"
echo ""

# ── Install Claude Code if needed ────────────────────────────────────

if ! command -v claude &> /dev/null; then
    if command -v npm &> /dev/null; then
        echo "Installing Claude Code CLI..."
        npm install -g @anthropic-ai/claude-code && \
            info "Claude Code CLI installed" || \
            warn "Could not install Claude Code CLI (try: npm install -g @anthropic-ai/claude-code)"
    else
        warn "Claude Code CLI not found (requires npm)"
        echo "  Install Node.js from https://nodejs.org then run:"
        echo "  npm install -g @anthropic-ai/claude-code"
        echo ""
    fi
fi

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

# ── Collect credentials ──────────────────────────────────────────────

echo ""
echo -e "${BOLD}Configure MCP credentials${NC}"
echo "Press Enter to skip any service you don't use."
echo ""

# Notion
echo -e "${BOLD}Notion${NC}"
prompt NOTION_API_KEY "API Key (ntn_...)"

# Slack
echo -e "${BOLD}Slack${NC}"
prompt SLACK_BOT_TOKEN "Bot Token (xoxb-...)"

# Clay
echo -e "${BOLD}Clay${NC}"
prompt CLAY_WORKSPACE_ID "Workspace ID"
prompt CLAY_SESSION_COOKIE "Session Cookie (s%3A...)"

# Namecheap
echo -e "${BOLD}Namecheap${NC}"
prompt NAMECHEAP_API_USER "API User"
prompt NAMECHEAP_API_KEY "API Key"
prompt NAMECHEAP_USERNAME "Username"
prompt NAMECHEAP_CLIENT_IP "Client IP"

# Premium Inboxes
echo -e "${BOLD}Premium Inboxes${NC}"
prompt PREMIUMINBOXES_API_TOKEN "API Token"

# ── Register MCPs with Claude Code (with credentials) ────────────────

if command -v claude &> /dev/null; then
    echo ""
    echo "Registering MCP servers with Claude Code..."

    # Only re-register if credentials were provided (skip preserves existing)
    register_claude_code() {
        local name="$1" binary="$2"
        shift 2
        local env_args=()
        local has_new_creds=false
        while [ $# -gt 0 ]; do
            local key="$1" val="$2"
            shift 2
            if [ -n "$val" ]; then
                env_args+=(-e "${key}=${val}")
                has_new_creds=true
            fi
        done
        if [ "$has_new_creds" = true ]; then
            claude mcp remove "$name" 2>/dev/null || true
            claude mcp add -s user "${env_args[@]}" "$name" -- "$binary" 2>/dev/null && \
                info "$name — registered with Claude Code" || \
                warn "$name — could not register"
        elif ! claude mcp get "$name" &>/dev/null; then
            claude mcp add -s user "$name" -- "$binary" 2>/dev/null && \
                info "$name — registered with Claude Code (no credentials)" || \
                warn "$name — could not register"
        else
            info "$name — already registered, keeping existing credentials"
        fi
    }

    # Notion (npx-based, not a local binary)
    register_claude_code_npx() {
        local name="$1"
        shift
        local env_args=()
        local has_new_creds=false
        while [ $# -gt 0 ]; do
            local key="$1" val="$2"
            shift 2
            if [ -n "$val" ]; then
                env_args+=(-e "${key}=${val}")
                has_new_creds=true
            fi
        done
        if [ "$has_new_creds" = true ]; then
            claude mcp remove "$name" 2>/dev/null || true
            claude mcp add -s user "${env_args[@]}" "$name" -- npx -y @modelcontextprotocol/server-notion 2>/dev/null && \
                info "$name — registered with Claude Code" || \
                warn "$name — could not register"
        elif ! claude mcp get "$name" &>/dev/null; then
            claude mcp add -s user "$name" -- npx -y @modelcontextprotocol/server-notion 2>/dev/null && \
                info "$name — registered with Claude Code (no credentials)" || \
                warn "$name — could not register"
        else
            info "$name — already registered, keeping existing credentials"
        fi
    }

    register_claude_code_npx notion \
        NOTION_API_KEY "$NOTION_API_KEY"

    register_claude_code slack "$CLAUDE_BIN_DIR/slack-mcp" \
        SLACK_BOT_TOKEN "$SLACK_BOT_TOKEN"

    register_claude_code clay "$CLAUDE_BIN_DIR/clay-mcp-server" \
        CLAY_WORKSPACE_ID "$CLAY_WORKSPACE_ID" \
        CLAY_SESSION_COOKIE "$CLAY_SESSION_COOKIE"

    register_claude_code namecheap "$CLAUDE_BIN_DIR/namecheap-mcp" \
        NAMECHEAP_API_USER "$NAMECHEAP_API_USER" \
        NAMECHEAP_API_KEY "$NAMECHEAP_API_KEY" \
        NAMECHEAP_USERNAME "$NAMECHEAP_USERNAME" \
        NAMECHEAP_CLIENT_IP "$NAMECHEAP_CLIENT_IP"

    register_claude_code premiuminboxes "$CLAUDE_BIN_DIR/premiuminboxes-mcp" \
        PREMIUMINBOXES_API_TOKEN "$PREMIUMINBOXES_API_TOKEN"
fi

# ── Configure Claude Desktop ─────────────────────────────────────────

CLAUDE_DESKTOP_DIR="$(dirname "$CLAUDE_DESKTOP_CONFIG")"
if [ -d "$CLAUDE_DESKTOP_DIR" ]; then
    echo ""
    echo "Configuring Claude Desktop..."
    python3 -c "
import json, os, sys

config_path = os.path.expanduser('~/Library/Application Support/Claude/claude_desktop_config.json')
config = {}
if os.path.exists(config_path):
    with open(config_path) as f:
        config = json.load(f)

bin_dir = os.path.expanduser('~/.claude/bin')
servers = config.setdefault('mcpServers', {})

# Credentials passed as args: key=value pairs
creds = {}
for arg in sys.argv[1:]:
    k, v = arg.split('=', 1)
    creds[k] = v

def env_or_existing(server_name, key):
    val = creds.get(key, '')
    if val:
        return val
    return servers.get(server_name, {}).get('env', {}).get(key, '')

mcp_defs = {
    'notion': {
        'command': 'npx',
        'args': ['-y', '@modelcontextprotocol/server-notion'],
        'env': {
            'NOTION_API_KEY': env_or_existing('notion', 'NOTION_API_KEY'),
        },
    },
    'slack': {
        'command': f'{bin_dir}/slack-mcp',
        'env': {
            'SLACK_BOT_TOKEN': env_or_existing('slack', 'SLACK_BOT_TOKEN'),
        },
    },
    'clay': {
        'command': f'{bin_dir}/clay-mcp-server',
        'env': {
            'CLAY_WORKSPACE_ID': env_or_existing('clay', 'CLAY_WORKSPACE_ID'),
            'CLAY_SESSION_COOKIE': env_or_existing('clay', 'CLAY_SESSION_COOKIE'),
        },
    },
    'namecheap': {
        'command': f'{bin_dir}/namecheap-mcp',
        'env': {
            'NAMECHEAP_API_USER': env_or_existing('namecheap', 'NAMECHEAP_API_USER'),
            'NAMECHEAP_API_KEY': env_or_existing('namecheap', 'NAMECHEAP_API_KEY'),
            'NAMECHEAP_USERNAME': env_or_existing('namecheap', 'NAMECHEAP_USERNAME'),
            'NAMECHEAP_CLIENT_IP': env_or_existing('namecheap', 'NAMECHEAP_CLIENT_IP'),
        },
    },
    'premiuminboxes': {
        'command': f'{bin_dir}/premiuminboxes-mcp',
        'env': {
            'PREMIUMINBOXES_API_TOKEN': env_or_existing('premiuminboxes', 'PREMIUMINBOXES_API_TOKEN'),
        },
    },
}

for name, defn in mcp_defs.items():
    servers[name] = defn

with open(config_path, 'w') as f:
    json.dump(config, f, indent=2)
" \
    "NOTION_API_KEY=$NOTION_API_KEY" \
    "SLACK_BOT_TOKEN=$SLACK_BOT_TOKEN" \
    "CLAY_WORKSPACE_ID=$CLAY_WORKSPACE_ID" \
    "CLAY_SESSION_COOKIE=$CLAY_SESSION_COOKIE" \
    "NAMECHEAP_API_USER=$NAMECHEAP_API_USER" \
    "NAMECHEAP_API_KEY=$NAMECHEAP_API_KEY" \
    "NAMECHEAP_USERNAME=$NAMECHEAP_USERNAME" \
    "NAMECHEAP_CLIENT_IP=$NAMECHEAP_CLIENT_IP" \
    "PREMIUMINBOXES_API_TOKEN=$PREMIUMINBOXES_API_TOKEN" \
    && info "Claude Desktop config updated" || warn "Could not update Claude Desktop config"
else
    warn "Claude Desktop not found — skipping config"
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
echo "MCP servers:"
[ -n "$NOTION_API_KEY" ]           && info "notion — configured"         || warn "notion — no credentials (configure later)"
[ -n "$SLACK_BOT_TOKEN" ]          && info "slack — configured"          || warn "slack — no credentials (configure later)"
[ -n "$CLAY_WORKSPACE_ID" ]        && info "clay — configured"           || warn "clay — no credentials (configure later)"
[ -n "$NAMECHEAP_API_KEY" ]        && info "namecheap — configured"      || warn "namecheap — no credentials (configure later)"
[ -n "$PREMIUMINBOXES_API_TOKEN" ] && info "premiuminboxes — configured" || warn "premiuminboxes — no credentials (configure later)"
echo ""
echo "To update credentials later, re-run this installer or edit:"
echo "  Claude Desktop: $CLAUDE_DESKTOP_CONFIG"
echo "  Claude Code:    claude mcp add -e KEY=val -s user <name> -- ~/.claude/bin/<binary>"
echo ""
echo "Restart Claude Desktop to activate MCP servers."
echo ""
