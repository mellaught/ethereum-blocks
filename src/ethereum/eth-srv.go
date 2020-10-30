package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// EthereumSRV ...
type EthereumSRV struct {
	logger *logrus.Logger    // logger
	cli    *ethclient.Client // ethereun client
	blocks *BlocksTxsStorage
}

// CreateNewEthereumSRV creates new ethereum block service
// Input: logger, ethereum explorer(or node) url
// Output: ethereum service
func CreateNewEthereumSRV(logger *logrus.Logger, URL string) *EthereumSRV {
	client, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}

	return &EthereumSRV{
		logger: logger,
		cli:    client,
		blocks: CreateNewBlocksTxsStorage(logger, client, URL),
	}
}
