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

func ConsumerrequestcountsallHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		consumer_idVal, ok := args["consumer_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: consumer_id"), nil
		}
		consumer_id, ok := consumer_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: consumer_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["start_datetime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_datetime=%v", val))
		}
		if val, ok := args["end_datetime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end_datetime=%v", val))
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
		url := fmt.Sprintf("%s/vault/consumers/%s/stats%s", cfg.BaseURL, consumer_id, queryString)
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
		var result models.ConsumerRequestCountsInDateRangeResponse
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

func CreateConsumerrequestcountsallTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_vault_consumers_consumer_id_stats",
		mcp.WithDescription("Consumer request counts"),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("consumer_id", mcp.Required(), mcp.Description("ID of the consumer to return")),
		mcp.WithString("start_datetime", mcp.Required(), mcp.Description("Scopes results to requests that happened after datetime")),
		mcp.WithString("end_datetime", mcp.Required(), mcp.Description("Scopes results to requests that happened before datetime")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ConsumerrequestcountsallHandler(cfg),
	}
}
