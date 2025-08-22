package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/vault-api/mcp-server/config"
	"github.com/vault-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func SessionscreateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Session
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/vault/sessions", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			req.Header.Set("code", cfg.APIKey)
		}
		if cfg.APIKey != "" {
			req.Header.Set("redirect_uri", cfg.APIKey)
		}
		if cfg.APIKey != "" {
			req.Header.Set("scope", cfg.APIKey)
		}
		if cfg.APIKey != "" {
			req.Header.Set("state", cfg.APIKey)
		}
		if cfg.BearerToken != "" {
			req.Header.Set("x-apideck-downstream-authorization", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-apideck-consumer-id"]; ok {
			req.Header.Set("x-apideck-consumer-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-app-id"]; ok {
			req.Header.Set("x-apideck-app-id", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateSessionResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateSessionscreateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_vault_sessions",
		mcp.WithDescription("Create Session"),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithObject("consumer_metadata", mcp.Description("Input parameter: The metadata of the consumer. This is used to display the consumer in the sidebar. This is optional, but recommended.")),
		mcp.WithObject("custom_consumer_settings", mcp.Description("Input parameter: Custom consumer settings that are passed as part of the session.")),
		mcp.WithString("redirect_uri", mcp.Description("Input parameter: The URL to redirect the user to after the session has been configured.")),
		mcp.WithObject("settings", mcp.Description("Input parameter: Settings to change the way the Vault is displayed.")),
		mcp.WithObject("theme", mcp.Description("Input parameter: Theming options to change the look and feel of Vault.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SessionscreateHandler(cfg),
	}
}
