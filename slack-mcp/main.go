package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/slack-go/slack"
)

var slackClient *slack.Client

func main() {
	// Initialize Slack client
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "SLACK_BOT_TOKEN environment variable required")
		os.Exit(1)
	}
	slackClient = slack.New(token)

	// Create MCP server
	s := server.NewMCPServer(
		"Slack MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Register tools
	registerTools(s)

	// Start stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func registerTools(s *server.MCPServer) {
	// Tool: Create Channel
	createChannelTool := mcp.NewTool("create_channel",
		mcp.WithDescription("Create a new Slack channel"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Channel name (lowercase, no spaces)"),
		),
		mcp.WithBoolean("is_private",
			mcp.Description("Whether the channel should be private"),
		),
	)
	s.AddTool(createChannelTool, createChannelHandler)

	// Tool: Invite Users
	inviteUsersTool := mcp.NewTool("invite_users",
		mcp.WithDescription("Invite users to a Slack channel"),
		mcp.WithString("channel_id",
			mcp.Required(),
			mcp.Description("The channel ID to invite users to"),
		),
		mcp.WithString("user_ids",
			mcp.Required(),
			mcp.Description("Comma-separated list of user IDs"),
		),
	)
	s.AddTool(inviteUsersTool, inviteUsersHandler)

	// Tool: Send Message
	sendMessageTool := mcp.NewTool("send_message",
		mcp.WithDescription("Send a message to a Slack channel"),
		mcp.WithString("channel_id",
			mcp.Required(),
			mcp.Description("The channel ID to send the message to"),
		),
		mcp.WithString("message",
			mcp.Required(),
			mcp.Description("The message text to send"),
		),
	)
	s.AddTool(sendMessageTool, sendMessageHandler)

	// Tool: Lookup User by Email
	lookupUserTool := mcp.NewTool("lookup_user",
		mcp.WithDescription("Look up a Slack user by email"),
		mcp.WithString("email",
			mcp.Required(),
			mcp.Description("The user's email address"),
		),
	)
	s.AddTool(lookupUserTool, lookupUserHandler)
}

func createChannelHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)
	isPrivate := false
	if v, ok := args["is_private"].(bool); ok {
		isPrivate = v
	}

	channel, err := slackClient.CreateConversation(slack.CreateConversationParams{
		ChannelName: name,
		IsPrivate:   isPrivate,
	})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create channel: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Channel created successfully!\nID: %s\nName: #%s",
		channel.ID, channel.Name,
	)), nil
}

func inviteUsersHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	channelID := args["channel_id"].(string)
	userIDs := args["user_ids"].(string)

	channel, err := slackClient.InviteUsersToConversation(channelID, userIDs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to invite users: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Users invited to #%s successfully!", channel.Name,
	)), nil
}

func sendMessageHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	channelID := args["channel_id"].(string)
	message := args["message"].(string)

	_, timestamp, err := slackClient.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to send message: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Message sent successfully! Timestamp: %s", timestamp,
	)), nil
}

func lookupUserHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	email := args["email"].(string)

	user, err := slackClient.GetUserByEmail(email)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to find user: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"User found!\nID: %s\nName: %s\nReal Name: %s",
		user.ID, user.Name, user.RealName,
	)), nil
}
