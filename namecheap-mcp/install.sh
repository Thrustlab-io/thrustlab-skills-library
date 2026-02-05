#!/bin/bash

# Namecheap MCP Server Installation Script

set -e

echo "🚀 Installing Namecheap MCP Server..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed. Please install Go 1.23.2 or later."
    exit 1
fi

echo "✅ Go found: $(go version)"

# Build the server
echo "📦 Building Namecheap MCP server..."
go mod download
go build -o namecheap-mcp

if [ ! -f "namecheap-mcp" ]; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"

# Create .env if it doesn't exist
if [ ! -f ".env" ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "⚠️  Please edit .env with your Namecheap API credentials"
fi

# Get the absolute path
INSTALL_PATH=$(pwd)/namecheap-mcp

echo ""
echo "✅ Installation complete!"
echo ""
echo "Binary location: $INSTALL_PATH"
echo ""
echo "📋 Next steps:"
echo "1. Edit .env with your Namecheap API credentials"
echo "2. Enable API access in your Namecheap account (Profile > Tools > API Access)"
echo "3. Whitelist your IP address in Namecheap API settings"
echo "4. Add to Claude Desktop configuration:"
echo ""
echo "Add this to ~/Library/Application Support/Claude/claude_desktop_config.json:"
echo ""
cat << EOF
{
  "mcpServers": {
    "namecheap": {
      "command": "$INSTALL_PATH",
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
EOF
echo ""
echo "5. Restart Claude Desktop"
echo ""
echo "Happy domain hunting! 🌐"
