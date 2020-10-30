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
