# Premium Inboxes MCP Server

A Model Context Protocol (MCP) server for Premium Inboxes API integration. This server provides tools to interact with Premium Inboxes subscriptions, orders, and purchases.

## Features

This MCP server provides the following tools:

### Subscription Management
- `get_all_subscriptions` - Get all subscriptions with orders, tags, billing dates, items, and pricing
- `get_subscription` - Get a single subscription by ID with complete details
- `cancel_subscription` - Cancel a subscription with optional reason and immediate removal

### Order Management
- `get_all_orders` - Get all orders with status, domains, inboxes, emails, and tags
- `get_order` - Get a single order by ID with complete details

### Purchase Management
- `create_purchase` - Create a new purchase order with hosting platform, sequencer details, domain info, email configuration, personas, and coupons

## Prerequisites

- Go 1.23.2 or later
- Premium Inboxes API token (obtain from your Premium Inboxes settings page)

## Installation

1. Clone this repository or navigate to the `premiuminboxes-mcp` directory

2. Install dependencies:
```bash
go mod download
```

3. Build the server:
```bash
go build -o premiuminboxes-mcp
```

## Configuration

Set your Premium Inboxes API token as an environment variable:

```bash
export PREMIUMINBOXES_API_TOKEN="your-api-token-here"
```

You can generate your API token from the Premium Inboxes Settings page at https://api.premiuminboxes.com/client/swagger

## Usage

### Running the Server

The MCP server communicates over stdio:

```bash
./premiuminboxes-mcp
```

### Using with Claude Desktop

Add this to your Claude Desktop configuration file:

**MacOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "premiuminboxes": {
      "command": "/path/to/premiuminboxes-mcp/premiuminboxes-mcp",
      "env": {
        "PREMIUMINBOXES_API_TOKEN": "your-api-token-here"
      }
    }
  }
}
```

### Example Tool Usage

#### Get All Subscriptions
```
Use the get_all_subscriptions tool to view all active subscriptions
```

#### Get Specific Subscription
```
Use the get_subscription tool with subscription_id: "sub_123456"
```

#### Cancel Subscription
```
Use the cancel_subscription tool with:
- subscription_id: "sub_123456"
- reason: "Customer requested cancellation"
- remove_immediately: false
```

#### Get All Orders
```
Use the get_all_orders tool to view all orders
```

#### Get Specific Order
```
Use the get_order tool with order_id: "order_123456"
```

#### Create Purchase
```
Use the create_purchase tool with purchase_data as a JSON string:
{
  "hostingPlatform": "aws",
  "sequencer": {...},
  "domain": {...},
  "emailConfig": {...},
  "personas": [...],
  "coupons": [...]
}
```

## API Authentication

All requests to the Premium Inboxes API require authentication via the `x-api-token` header. The server automatically includes this header using the token from the `PREMIUMINBOXES_API_TOKEN` environment variable.

## Error Handling

The server provides detailed error messages for:
- Missing API token
- Network errors
- API errors (with status codes)
- Invalid JSON in requests/responses

## Development

### Project Structure

```
premiuminboxes-mcp/
├── main.go          # Main server implementation
├── go.mod           # Go module definition
├── go.sum           # Go dependencies checksums
└── README.md        # This file
```

### Dependencies

- `github.com/mark3labs/mcp-go` - MCP protocol implementation
- Standard library only for HTTP client (`net/http`)

### Building from Source

```bash
cd premiuminboxes-mcp
go mod download
go build -o premiuminboxes-mcp
```

## License

See the main project license.

## Support

For issues with the Premium Inboxes API, visit: https://api.premiuminboxes.com/client/swagger

For MCP server issues, please file an issue in the repository.
