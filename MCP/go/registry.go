package main

import (
	"github.com/background-removal-api/mcp-server/config"
	"github.com/background-removal-api/mcp-server/models"
	tools_background_removal "github.com/background-removal-api/mcp-server/tools/background_removal"
	tools_fetch_account_info "github.com/background-removal-api/mcp-server/tools/fetch_account_info"
	tools_improvement_program "github.com/background-removal-api/mcp-server/tools/improvement_program"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_background_removal.CreatePost_removebgTool(cfg),
		tools_fetch_account_info.CreateGet_accountTool(cfg),
		tools_improvement_program.CreatePost_improveTool(cfg),
	}
}
