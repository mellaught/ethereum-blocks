package config

import (
	"github.com/mellaught/ethereum-blocks/src/models"
)

// Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() *models.ServiceConfig {
	return &models.ServiceConfig{
		Host: v.GetString("service.host"),
		Port: v.GetString("service.port"),
	}
}

// Reads explorer params from config.json
func (v *viperConfig) ReadExplorerConfig() *models.EthereumExplorer {
	return &models.EthereumExplorer{
		URL: v.GetString("infura.url"),
	}
}
