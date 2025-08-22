package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/vault-api/mcp-server/config"
	"github.com/vault-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ConnectionsupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		service_idVal, ok := args["service_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: service_id"), nil
		}
		service_id, ok := service_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: service_id"), nil
		}
		unified_apiVal, ok := args["unified_api"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: unified_api"), nil
		}
		unified_api, ok := unified_apiVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: unified_api"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("code=%s", cfg.APIKey))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("redirect_uri=%s", cfg.APIKey))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("scope=%s", cfg.APIKey))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("state=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("x-apideck-downstream-authorization=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Connection
		
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
		url := fmt.Sprintf("%s/vault/connections/%s/%s%s", cfg.BaseURL, service_id, unified_api, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
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
		var result models.UpdateConnectionResponse
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

func CreateConnectionsupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_vault_connections_unified_api_service_id",
		mcp.WithDescription("Update connection"),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("service_id", mcp.Required(), mcp.Description("Service ID of the resource to return")),
		mcp.WithString("unified_api", mcp.Required(), mcp.Description("Unified API")),
		mcp.WithString("auth_type", mcp.Description("Input parameter: Type of authorization used by the connector")),
		mcp.WithString("authorize_url", mcp.Description("Input parameter: The OAuth redirect URI. Redirect your users to this URI to let them authorize your app in the connector's UI. Before you can use this URI, you must add `redirect_uri` as a query parameter to the `authorize_url`. Be sure to URL encode the `redirect_uri` part. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique identifier of the connection.")),
		mcp.WithArray("resource_schema_support", mcp.Description("")),
		mcp.WithString("website", mcp.Description("Input parameter: The website URL of the connection")),
		mcp.WithString("icon", mcp.Description("Input parameter: A visual icon of the connection, that will be shown in the Vault")),
		mcp.WithString("revoke_url", mcp.Description("Input parameter: The OAuth revoke URI. Redirect your users to this URI to revoke this connection. Before you can use this URI, you must add `redirect_uri` as a query parameter. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.")),
		mcp.WithString("state", mcp.Description("Input parameter: [Connection state flow](#section/Connection-state)")),
		mcp.WithArray("form_fields", mcp.Description("Input parameter: The settings that are wanted to create a connection.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the connection")),
		mcp.WithBoolean("enabled", mcp.Description("Input parameter: Whether the connection is enabled or not. You can enable or disable a connection using the Update Connection API.")),
		mcp.WithObject("metadata", mcp.Description("Input parameter: Attach your own consumer specific metadata")),
		mcp.WithString("oauth_grant_type", mcp.Description("Input parameter: OAuth grant type used by the connector. More info: https://oauth.net/2/grant-types")),
		mcp.WithObject("settings", mcp.Description("Input parameter: Connection settings. Values will persist to `form_fields` with corresponding id")),
		mcp.WithString("tag_line", mcp.Description("")),
		mcp.WithArray("subscriptions", mcp.Description("")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of the connection.")),
		mcp.WithBoolean("schema_support", mcp.Description("")),
		mcp.WithString("unified_api", mcp.Description("Input parameter: The unified API category where the connection belongs to.")),
		mcp.WithBoolean("has_guide", mcp.Description("Input parameter: Whether the connector has a guide available in the developer docs or not (https://docs.apideck.com/connectors/{service_id}/docs/consumer+connection).")),
		mcp.WithString("integration_state", mcp.Description("Input parameter: The current state of the Integration.")),
		mcp.WithString("service_id", mcp.Description("Input parameter: The ID of the service this connection belongs to.")),
		mcp.WithBoolean("validation_support", mcp.Description("")),
		mcp.WithArray("configuration", mcp.Description("")),
		mcp.WithArray("configurable_resources", mcp.Description("")),
		mcp.WithArray("resource_settings_support", mcp.Description("")),
		mcp.WithString("created_at", mcp.Description("")),
		mcp.WithString("updated_at", mcp.Description("")),
		mcp.WithArray("custom_mappings", mcp.Description("Input parameter: List of custom mappings configured for this connection")),
		mcp.WithString("logo", mcp.Description("Input parameter: The logo of the connection, that will be shown in the Vault")),
		mcp.WithArray("settings_required_for_authorization", mcp.Description("Input parameter: List of settings that are required to be configured on integration before authorization can occur")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ConnectionsupdateHandler(cfg),
	}
}
