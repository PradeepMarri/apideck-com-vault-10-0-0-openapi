package main

import (
	"github.com/vault-api/mcp-server/config"
	"github.com/vault-api/mcp-server/models"
	tools_connections "github.com/vault-api/mcp-server/tools/connections"
	tools_consumers "github.com/vault-api/mcp-server/tools/consumers"
	tools_logs "github.com/vault-api/mcp-server/tools/logs"
	tools_custom_mappings "github.com/vault-api/mcp-server/tools/custom_mappings"
	tools_sessions "github.com/vault-api/mcp-server/tools/sessions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_connections.CreateConnectionsimportTool(cfg),
		tools_consumers.CreateConsumersallTool(cfg),
		tools_consumers.CreateConsumersaddTool(cfg),
		tools_consumers.CreateConsumerrequestcountsallTool(cfg),
		tools_logs.CreateLogsallTool(cfg),
		tools_custom_mappings.CreateCustommappingsdeleteTool(cfg),
		tools_custom_mappings.CreateCustommappingsoneTool(cfg),
		tools_custom_mappings.CreateCustommappingsupdateTool(cfg),
		tools_custom_mappings.CreateCustommappingsaddTool(cfg),
		tools_connections.CreateCustomfieldsallTool(cfg),
		tools_consumers.CreateConsumersupdateTool(cfg),
		tools_consumers.CreateConsumersdeleteTool(cfg),
		tools_consumers.CreateConsumersoneTool(cfg),
		tools_connections.CreateConnectionsexampleTool(cfg),
		tools_connections.CreateConnectionsschemaTool(cfg),
		tools_connections.CreateConnectionstokenTool(cfg),
		tools_connections.CreateConnectionsrevokeTool(cfg),
		tools_connections.CreateConnectionsaddTool(cfg),
		tools_connections.CreateConnectionsdeleteTool(cfg),
		tools_connections.CreateConnectionsoneTool(cfg),
		tools_connections.CreateConnectionsupdateTool(cfg),
		tools_connections.CreateConnectionsettingsallTool(cfg),
		tools_connections.CreateConnectionsettingsupdateTool(cfg),
		tools_sessions.CreateSessionscreateTool(cfg),
		tools_connections.CreateConnectionscallbackTool(cfg),
		tools_connections.CreateConnectionsauthorizeTool(cfg),
		tools_connections.CreateConnectionsallTool(cfg),
	}
}
