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

type Block struct {
	Number     uint64         `json:"header"`
	ParentHash string         `json:"parent_hash"`
	Difficulty string         `json:"difficulty"`
	TxCount    int            `json:"tx_count"`
	Txs        []*Transaction `json:"transactions"`
	Timestamp  uint64         `json:"timestamp"`
}

type Transaction struct {
	Hash  string `json:"hash"`
	To    string `json:"to"`
	Value string `json:"value"`
	Nonce uint64 `json:"nonce"`
}
