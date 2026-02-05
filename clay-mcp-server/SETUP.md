# Clay MCP Server - Quick Setup Guide

## Installation via Install Script

### 1. Run the Install Script

From the thrustlab directory:

```bash
./install.sh
```

This will:
- Install all Thrustlab skills
- Build and install Clay MCP to `~/.claude/bin/clay-mcp-server`
- Register with Claude (if credentials are set)

### 2. Get Your Clay Credentials

#### A. Get Workspace ID

1. Open Clay.com in your browser
2. Navigate to any workbook
3. Look at the URL: `https://app.clay.com/workspaces/757984/workbooks/...`
4. Copy the number after `/workspaces/` (e.g., `757984`)

#### B. Get Session Cookie

**Option 1: Firefox DevTools** (Recommended)

1. Open Clay.com in Firefox
2. Press `F12` → **Storage** tab → **Cookies**
3. Find `claysession` cookie
4. Copy the full **Value** (starts with `s%3A...`)

**Option 2: From Network HAR**

1. Press `F12` → **Network** tab
2. Enable "Persist Logs"
3. Reload page
4. Right-click → "Save All As HAR"
5. Search file for `"claysession"`
6. Copy the value after `claysession=`

#### C. Get Workbook ID (for creating tables)

1. Open the workbook where you want to create tables
2. Look at URL: `https://app.clay.com/workspaces/757984/workbooks/wb_abc123/...`
3. Copy the workbook ID (e.g., `wb_abc123`)

## Configuration Methods

### Method 1: Chat Commands (Recommended - Easiest!)

Run the install script first:

```bash
./install.sh
```

Then configure via chat - no manual file editing needed:

**In Claude:**

```
Set my Clay workspace ID to 757984
```

```
Set my Clay session cookie to s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx...
```

Restart Claude Desktop and you're done! ✨

### Method 2: Environment Variables (Advanced)

Register with credentials during installation:

```bash
export CLAY_WORKSPACE_ID="757984"
export CLAY_SESSION_COOKIE="s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx..."

claude mcp add \
  -e CLAY_WORKSPACE_ID="$CLAY_WORKSPACE_ID" \
  -e CLAY_SESSION_COOKIE="$CLAY_SESSION_COOKIE" \
  -s user clay -- ~/.claude/bin/clay-mcp-server
```

### Method 3: Manual Config Edit (For Reference)

After running the install script without credentials, configure via chat:

**In Claude:**

```
Set my Clay workspace ID to 757984
```

Then:

```
Set my Clay session cookie to s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx...
```

**What happens:**
- Updates your Claude Desktop config file automatically
- Changes are permanent (written to `~/Library/Application Support/Claude/claude_desktop_config.json`)
- Restart Claude Desktop to activate the new settings

**Note:** This is the easiest setup method - just chat to configure!

## Verify Installation

Once configured, test it:

```
Create a Clay table to find accounting firms in Belgium
with 50 employees. Use workbook wb_0t9zx9hiFdeR4VDUCHj.
```

Claude should use the `search_companies_by_industry` tool and create a table.

## Usage Examples

### Industry Search

```
Find software companies in France with 200+ employees
that mention "SaaS" in workbook wb_xxx
```

### Geography Search

```
Find all restaurants within 5km of Brussels (50.8503, 4.3517)
in workbook wb_yyy
```

### Update Configuration

```
Update my Clay workspace ID to 123456
```

## Troubleshooting

### "Authentication failed"

Your session cookie expired. Get a fresh one:
1. Open Clay.com
2. Get new cookie from DevTools
3. Run: `Set my Clay session cookie to <new-cookie>`

### "Workspace not found"

Check your workspace ID in the Clay.com URL.

### "MCP not showing up"

1. Restart Claude Desktop
2. Check `claude mcp list` shows "clay"
3. Try `claude mcp get clay` for details

## Security Notes

- **Never share your session cookie** - it grants full access to your Clay account
- **Don't commit to git** - keep credentials out of version control
- **Rotate regularly** - get fresh cookies periodically
- Session cookies expire after some time - you'll need to update them

## Manual Installation

If not using the install script:

```bash
cd /Users/kwinten/thrustlab/clay-mcp-server
go build -o clay-mcp-server

# Copy to Claude bin directory
mkdir -p ~/.claude/bin
cp clay-mcp-server ~/.claude/bin/

# Register with Claude
claude mcp add \
  -e CLAY_WORKSPACE_ID="your-id" \
  -e CLAY_SESSION_COOKIE="your-cookie" \
  -s user clay -- ~/.claude/bin/clay-mcp-server
```

## Next Steps

Once configured, you can:
- Create company search tables programmatically
- Create geography-based business searches
- Iterate on Clay prompts from Claude chat
- Automate GTM campaign setup

See [README.md](README.md) for detailed API documentation and examples.
