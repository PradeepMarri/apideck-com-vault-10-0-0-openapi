package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vault-api/mcp-server/config"
	"github.com/vault-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ConnectionsettingsallHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		unified_apiVal, ok := args["unified_api"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: unified_api"), nil
		}
		unified_api, ok := unified_apiVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: unified_api"), nil
		}
		service_idVal, ok := args["service_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: service_id"), nil
		}
		service_id, ok := service_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: service_id"), nil
		}
		resourceVal, ok := args["resource"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resource"), nil
		}
		resource, ok := resourceVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resource"), nil
		}
		url := fmt.Sprintf("%s/vault/connections/%s/%s/%s/config", cfg.BaseURL, unified_api, service_id, resource)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetConnectionResponse
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

func CreateConnectionsettingsallTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_vault_connections_unified_api_service_id_resource_config",
		mcp.WithDescription("Get resource settings"),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("unified_api", mcp.Required(), mcp.Description("Unified API")),
		mcp.WithString("service_id", mcp.Required(), mcp.Description("Service ID of the resource to return")),
		mcp.WithString("resource", mcp.Required(), mcp.Description("Name of the resource (plural)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ConnectionsettingsallHandler(cfg),
	}
}
