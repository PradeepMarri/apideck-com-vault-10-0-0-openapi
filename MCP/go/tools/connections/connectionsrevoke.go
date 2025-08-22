package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vault-api/mcp-server/config"
	"github.com/vault-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ConnectionsrevokeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		application_idVal, ok := args["application_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: application_id"), nil
		}
		application_id, ok := application_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: application_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["redirect_uri"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("redirect_uri=%v", val))
		}
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
		url := fmt.Sprintf("%s/vault/revoke/%s/%s%s", cfg.BaseURL, service_id, application_id, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.UnexpectedErrorResponse
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

func CreateConnectionsrevokeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_vault_revoke_service_id_application_id",
		mcp.WithDescription("Revoke connection"),
		mcp.WithString("service_id", mcp.Required(), mcp.Description("Service ID of the resource to return")),
		mcp.WithString("application_id", mcp.Required(), mcp.Description("Application ID of the resource to return")),
		mcp.WithString("state", mcp.Required(), mcp.Description("An opaque value the applications adds to the initial request that the authorization server includes when redirecting the back to the application. This value must be used by the application to prevent CSRF attacks.")),
		mcp.WithString("redirect_uri", mcp.Required(), mcp.Description("URL to redirect back to after authorization. When left empty the default configured redirect uri will be used.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ConnectionsrevokeHandler(cfg),
	}
}
