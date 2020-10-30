package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/mellaught/ethereum-blocks/src/models"

	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	ReadServiceConfig() *models.ServiceConfig
	ReadExplorerConfig() *models.EthereumExplorer
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetFloat64(key string) float64
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	if _, err := os.Stat("./config.json.local"); !os.IsNotExist(err) {
		viper.SetConfigFile(`config.json.local`)
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// NewViperConfig creates new viper for reading config.json
func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
