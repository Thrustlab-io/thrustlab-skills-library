# Namecheap MCP Server

A Model Context Protocol (MCP) server for interacting with the Namecheap domain registration API.

## Features

- **Check Domain Availability**: Search for available domains
- **Register Domains**: Create/register new domain names
- **List Domains**: View all domains in your account

## Prerequisites

- Go 1.23.2 or later
- Namecheap account with API access enabled
- Namecheap API credentials

## Setup

### 1. Enable Namecheap API Access

1. Log in to your Namecheap account
2. Go to Profile > Tools
3. Scroll to "Business & Dev Tools"
4. Click "MANAGE" next to "Namecheap API Access"
5. Toggle API access ON
6. Read and accept the Terms of Service
7. Whitelist your IP address

### 2. Get Your API Credentials

You'll need:
- **API User**: Your Namecheap username
- **API Key**: Generated in the API settings
- **Username**: Your Namecheap username
- **Client IP**: Your whitelisted IP address

### 3. Configure Environment Variables

Create a `.env` file or export these variables:

```bash
export NAMECHEAP_API_USER="your_username"
export NAMECHEAP_API_KEY="your_api_key"
export NAMECHEAP_USERNAME="your_username"
export NAMECHEAP_CLIENT_IP="your.ip.address"
export NAMECHEAP_SANDBOX="true"  # Set to "false" for production
```

### 4. Build the Server

```bash
cd namecheap-mcp
go mod download
go build -o namecheap-mcp
```

### 5. Run the Server

```bash
./namecheap-mcp
```

## Available Tools

### check_domain

Check if one or more domains are available for registration.

**Parameters:**
- `domains` (required): Comma-separated list of domain names (e.g., "example.com,mysite.net")

**Example:**
```json
{
  "domains": "example.com,mysite.net,coolsite.io"
}
```

### create_domain

Register a new domain name.

**Parameters:**
- `domain` (required): Domain name to register
- `years` (optional): Number of years (default: 1)
- `registrant_first_name` (required): First name
- `registrant_last_name` (required): Last name
- `registrant_address` (required): Street address
- `registrant_city` (required): City
- `registrant_state` (required): State/Province
- `registrant_postal_code` (required): Postal code
- `registrant_country` (required): Country code (e.g., "US")
- `registrant_phone` (required): Phone number
- `registrant_email` (required): Email address

**Example:**
```json
{
  "domain": "example.com",
  "years": 2,
  "registrant_first_name": "John",
  "registrant_last_name": "Doe",
  "registrant_address": "123 Main St",
  "registrant_city": "New York",
  "registrant_state": "NY",
  "registrant_postal_code": "10001",
  "registrant_country": "US",
  "registrant_phone": "+1.2125551234",
  "registrant_email": "john@example.com"
}
```

### list_domains

List all domains in your Namecheap account.

**Parameters:**
- `page` (optional): Page number (default: 1)
- `page_size` (optional): Domains per page (default: 20, max: 100)

**Example:**
```json
{
  "page": 1,
  "page_size": 50
}
```

## Testing

Use the sandbox environment for testing:

```bash
export NAMECHEAP_SANDBOX="true"
```

Sandbox API URL: `https://api.sandbox.namecheap.com/xml.response`

## Dependencies

- `github.com/mark3labs/mcp-go` - MCP SDK for Go
- Standard library only (`net/http`, `encoding/xml`)

## API Documentation

For more information about the Namecheap API:
- [API Methods](https://www.namecheap.com/support/api/methods/)
- [API Introduction](https://www.namecheap.com/support/api/intro/)
- [Global Parameters](https://www.namecheap.com/support/api/global-parameters/)

## Integration with Claude Desktop

Add to your Claude Desktop MCP configuration (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "namecheap": {
      "command": "/path/to/namecheap-mcp/namecheap-mcp",
      "env": {
        "NAMECHEAP_API_USER": "your_username",
        "NAMECHEAP_API_KEY": "your_api_key",
        "NAMECHEAP_USERNAME": "your_username",
        "NAMECHEAP_CLIENT_IP": "your.ip.address",
        "NAMECHEAP_SANDBOX": "true"
      }
    }
  }
}
```

## License

MIT
