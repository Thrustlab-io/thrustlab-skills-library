# Clay MCP Server - Workspace Switching Guide

## Overview

The Clay MCP server now supports easy switching between multiple Clay workspaces without requiring restarts. This is useful when working with multiple clients or managing multiple Clay accounts.

## Features

- **Workspace Profiles**: Save multiple workspace configurations with friendly names
- **Runtime Switching**: Switch between workspaces instantly without restarting Claude Desktop
- **Default Workspace**: Set a default workspace for automatic initialization
- **Backward Compatible**: Still supports environment variable configuration

## Quick Start

### 1. Add a Workspace Profile

```
add_workspace_profile(
  name='thrustlab',
  workspace_id='757984',
  session_cookie='s%3A...',
  description='Thrustlab main workspace',
  set_as_default=true
)
```

### 2. Add More Workspaces

```
add_workspace_profile(
  name='client-acme',
  workspace_id='123456',
  session_cookie='s%3A...',
  description='ACME Corp client workspace'
)
```

### 3. List Available Workspaces

```
list_workspace_profiles()
```

Output shows:
- All saved profiles
- Which one is the default
- Which one is currently active

### 4. Switch Between Workspaces

```
switch_workspace(name='client-acme')
```

**No restart required!** The switch is effective immediately for all subsequent operations.

### 5. Check Current Workspace

```
get_current_workspace()
```

Shows your currently active workspace ID and workbook.

## How It Works

### Initialization Priority

When the MCP server starts, it looks for credentials in this order:

1. **Environment Variables** (highest priority)
   - `CLAY_WORKSPACE_ID`
   - `CLAY_SESSION_COOKIE`
   - `CLAY_FRONTEND_VERSION`

2. **Default Workspace Profile** (fallback)
   - Loads from `~/.clay/workspaces.json`
   - Uses the profile marked as default

If neither is found, the server will show an error with setup instructions.

### Storage Location

Workspace profiles are stored in:
```
~/.clay/workspaces.json
```

This file contains your workspace IDs and session cookies. It's set with restrictive permissions (0600) for security.

### Profile File Structure

```json
{
  "workspaces": {
    "thrustlab": {
      "workspace_id": "757984",
      "session_cookie": "s%3A...",
      "description": "Thrustlab main workspace",
      "frontend_version": ""
    },
    "client-acme": {
      "workspace_id": "123456",
      "session_cookie": "s%3A...",
      "description": "ACME Corp client workspace"
    }
  },
  "default": "thrustlab"
}
```

## Available Tools

### Workspace Management

| Tool | Description | Restart Required? |
|------|-------------|-------------------|
| `add_workspace_profile` | Save a workspace profile | No |
| `list_workspace_profiles` | Show all saved profiles | No |
| `switch_workspace` | Switch to a different workspace | **No** ✨ |
| `get_current_workspace` | Show current workspace info | No |
| `remove_workspace_profile` | Delete a saved profile | No |

### Legacy Tools (Still Supported)

| Tool | Description | Restart Required? |
|------|-------------|-------------------|
| `set_workspace_id` | Update workspace ID in config | **Yes** |
| `set_session_cookie` | Update session cookie in config | **Yes** |

## Common Workflows

### Working with Multiple Clients

```bash
# Set up client workspaces once
add_workspace_profile(name='client-a', workspace_id='111111', session_cookie='...')
add_workspace_profile(name='client-b', workspace_id='222222', session_cookie='...')
add_workspace_profile(name='client-c', workspace_id='333333', session_cookie='...')

# Switch as needed during your work
switch_workspace(name='client-a')
create_workbook(name='Q1 2024 Campaign')
search_companies_by_industry(...)

switch_workspace(name='client-b')
create_workbook(name='Product Launch')
search_businesses_by_geography(...)
```

### Setting a New Default

```bash
# Method 1: When adding a profile
add_workspace_profile(
  name='new-main',
  workspace_id='...',
  session_cookie='...',
  set_as_default=true
)

# Method 2: Update existing profile to be default
# (You'll need to use add_workspace_profile with set_as_default=true on an existing profile)
```

### Removing Old Workspaces

```bash
# Remove a workspace profile
remove_workspace_profile(name='old-client')

# List to verify
list_workspace_profiles()
```

## Migration from Environment Variables

If you're currently using environment variables, you can migrate to profiles:

1. **Add your current workspace as a profile:**
   ```bash
   add_workspace_profile(
     name='main',
     workspace_id='<your current CLAY_WORKSPACE_ID>',
     session_cookie='<your current CLAY_SESSION_COOKIE>',
     set_as_default=true
   )
   ```

2. **Optional: Remove env vars from Claude config**

   You can now remove `CLAY_WORKSPACE_ID` and `CLAY_SESSION_COOKIE` from your `claude_desktop_config.json` if you prefer to use profiles exclusively.

3. **Restart Claude Desktop**

   The server will now initialize from your default profile instead of env vars.

## Security Notes

- Session cookies are sensitive credentials - treat them like passwords
- The `~/.clay/workspaces.json` file is created with restrictive permissions (0600)
- Session cookies are masked in output (only first/last 10 chars shown)
- Don't commit the workspaces.json file to version control
- Session cookies expire - you'll need to update them periodically

## Troubleshooting

### "No default workspace profile set"

Either:
- Set a default: `add_workspace_profile(..., set_as_default=true)`
- Or use environment variables instead

### "Profile 'xyz' does not exist"

Check available profiles with `list_workspace_profiles()` and use the exact name.

### Session expired errors

Your session cookie has expired. Get a new one from Clay.com and update the profile:

```bash
add_workspace_profile(
  name='existing-profile',
  workspace_id='...',
  session_cookie='<new cookie>',  # Updated cookie
  description='...'
)
```

## Advanced: Optional Per-Tool Workspace Override

While not currently implemented, the architecture supports adding an optional `workspace_id` parameter to individual tools if needed. This would allow one-off operations in a different workspace without switching the global context.

If you need this feature, please open an issue!
