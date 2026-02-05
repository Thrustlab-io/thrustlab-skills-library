# Clay MCP Server

Model Context Protocol (MCP) server for Clay.com - enables programmatic table creation and campaign setup from Claude.

## Features

- **Search Companies by Industry** - Create tables with companies from Mixrank/LinkedIn data
- **Search Businesses by Geography** - Create tables with local businesses from Google Maps
- **Iterate on Clay Prompts** - Modify search parameters programmatically
- **Full Session Authentication** - Uses your Clay session for API access

## Installation

### 1. Clone or Download

```bash
cd /Users/kwinten/thrustlab/clay-mcp-server
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Server

```bash
go build -o clay-mcp-server
```

## Configuration

### Get Your Clay Credentials

#### 1. Get Workspace ID

- Open Clay.com in your browser
- Navigate to any workbook
- Look at the URL: `https://app.clay.com/workspaces/757984/workbooks/...`
- Copy the workspace ID (e.g., `757984`)

#### 2. Get Session Cookie

**Method 1: From Browser DevTools**
1. Open Clay.com in Firefox
2. Press `F12` to open DevTools
3. Go to **Storage** tab → **Cookies**
4. Find `claysession` cookie
5. Copy the **Value** (starts with `s%3A...`)

**Method 2: From HAR File**
1. Export HAR file using the tools in `/private/tmp/.../scratchpad/`
2. Search for `"Cookie":` in the HAR file
3. Extract the `claysession=...` value

#### 3. Get Workbook ID (for creating tables)

- Open the workbook where you want to create tables
- Look at URL: `https://app.clay.com/workspaces/757984/workbooks/wb_0t9zx9hiFdeR4VDUCHj/...`
- Copy the workbook ID (e.g., `wb_0t9zx9hiFdeR4VDUCHj`)

### Set Environment Variables

```bash
export CLAY_WORKSPACE_ID="757984"
export CLAY_SESSION_COOKIE="s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx.AnKH%2BaF5sq7cgaQ7yPCwFOSHyPdV2Lq2LY4OH1kssEY"
export CLAY_FRONTEND_VERSION="v20260205_151537Z_19e7945c5e" # Optional
```

## Usage with Claude Desktop

### Add to Claude Desktop Config

Edit `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "clay": {
      "command": "/Users/kwinten/thrustlab/clay-mcp-server/clay-mcp-server",
      "env": {
        "CLAY_WORKSPACE_ID": "757984",
        "CLAY_SESSION_COOKIE": "s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx.AnKH...",
        "CLAY_FRONTEND_VERSION": "v20260205_151537Z_19e7945c5e"
      }
    }
  }
}
```

### Restart Claude Desktop

After updating the config, restart Claude Desktop to load the MCP server.

## Available Tools

### Configuration Tools

#### set_workspace_id

Set your Clay workspace ID by updating the Claude Desktop config file.

**Parameters:**
- `workspace_id` (required) - Your Clay workspace ID (e.g., "757984")

**Example:**
```
Set my Clay workspace ID to 757984
```

**What it does:**
- Updates `~/Library/Application Support/Claude/claude_desktop_config.json`
- Adds/updates `CLAY_WORKSPACE_ID` in the clay MCP environment variables
- Requires Claude Desktop restart to take effect

#### set_session_cookie

Set your Clay session cookie by updating the Claude Desktop config file.

**Parameters:**
- `session_cookie` (required) - Session cookie value from browser

**Example:**
```
Set my Clay session cookie to s%3AyyBobWOFnTZnRlhw7cJhRG1ncBqKUUmx...
```

**What it does:**
- Updates `~/Library/Application Support/Claude/claude_desktop_config.json`
- Adds/updates `CLAY_SESSION_COOKIE` in the clay MCP environment variables
- Requires Claude Desktop restart to take effect

**Note:** These tools provide a convenient way to configure credentials via chat. Changes are permanent (written to config file) but require restarting Claude Desktop to take effect.

### Search Tools

#### 1. search_companies_by_industry

Search for companies using Clay's Mixrank/LinkedIn data source.

**Parameters:**
- `workbook_id` (required) - Workbook ID to create table in
- `industries` (required) - Comma-separated industries (e.g., "Accounting,Consulting")
- `countries` - Comma-separated countries (e.g., "Belgium,Netherlands")
- `company_sizes` - Sizes: `1`, `2-10`, `11-50`, `50`, `200`, `500`, `1000`, `5000`, `10000`
- `keywords` - Description keywords to filter by
- `limit` - Max results (default: 25000)

**Example Usage in Claude:**

```
Create a Clay table to find accounting firms in Belgium with 50 employees
that mention "boekhouding" in their description. Use workbook wb_0t9zx9hiFdeR4VDUCHj.
```

### 2. search_businesses_by_geography

Search for local businesses using Google Maps.

**Parameters:**
- `workbook_id` (required) - Workbook ID to create table in
- `latitude` (required) - Latitude coordinate (e.g., 51.049)
- `longitude` (required) - Longitude coordinate (e.g., 3.725)
- `proximity_km` (required) - Search radius in km
- `business_types` (required) - Comma-separated types (e.g., "art_gallery,restaurant")
- `table_name` - Custom table name (optional)
- `table_emoji` - Table emoji (optional, default: 🌙)

**Example Usage in Claude:**

```
Find art galleries within 13km of Ghent, Belgium (lat: 51.049, lng: 3.725)
and create a table in workbook wb_0t9zx9hiFdeR4VDUCHj.
```

## Examples

### Industry Search

```
Hey Claude, I need to find software companies in France with 200+ employees
that mention "SaaS" in their description. Create the table in workbook wb_xxx.
```

Claude will use: `search_companies_by_industry`
- industries: "Software"
- countries: "France"
- company_sizes: "200,500,1000,5000,10000"
- keywords: "SaaS"

### Geography Search

```
Find all restaurants and cafes in a 5km radius around Brussels city center
(50.8503° N, 4.3517° E). Use workbook wb_yyy.
```

Claude will use: `search_businesses_by_geography`
- latitude: 50.8503
- longitude: 4.3517
- proximity_km: 5
- business_types: "restaurant,cafe"

### Iterating on Searches

```
Actually, let's expand that to 10km and also include bars.
```

Claude will call the tool again with updated parameters.

## API Structure

Based on reverse-engineered Clay.com API:

### Industry Search Endpoint
```
POST https://api.clay.com/v3/workspaces/{workspace_id}/wizard/evaluate-step
```

Uses Action Package: `e251a70e-46d7-4f3a-b3ef-a211ad3d8bd2` (Mixrank)

### Geography Search Endpoint
```
POST https://api.clay.com/v3/tables
```

Uses Action Package: `3282a1c7-6bb0-497e-a34b-32268e104e55` (Google Maps)

## Troubleshooting

### Session Expired

If you get authentication errors:
1. Open Clay.com in browser
2. Get a fresh session cookie
3. Update `CLAY_SESSION_COOKIE` environment variable
4. Restart Claude Desktop

### Workspace ID Invalid

Make sure you're using your actual workspace ID from the Clay.com URL.

### Table Not Appearing

Check the Clay web interface - tables may take a few seconds to populate with data.

## Development

### Run in Development Mode

```bash
go run main.go
```

### Test with MCP Inspector

```bash
npx @modelcontextprotocol/inspector clay-mcp-server
```

## Security Notes

- **Session Cookie Security**: Your session cookie grants full access to your Clay account. Keep it secure!
- **Don't commit credentials**: Never commit environment variables to git
- **Rotate regularly**: Clay sessions may expire - update your cookie periodically

## License

MIT

## Credits

Built by Thrustlab for programmatic GTM campaign creation.

Based on HAR analysis of Clay.com API endpoints.
