package models

// ServiceConfig contains configurations for rest-api service.
type ServiceConfig struct {
	Host string // Service Host
	Port string // Service port
}

// EthereumExplorer contains configurations about one of ethereum explorers
type EthereumExplorer struct {
	URL string
}
