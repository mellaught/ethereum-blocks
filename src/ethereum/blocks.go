package eth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// Blocks ...
type BlocksTxsStorage struct {
	sync.RWMutex
	blocks []*types.Block // array(size == 100) of last 100 blocks
	txs    map[string]*types.Transaction
	logger *logrus.Logger // logger
}

// CreateNewInfuraApp creates new infura application
// Input:
// Output:
func CreateNewBlocksTxsStorage(logger *logrus.Logger, cli *ethclient.Client, URL string) *BlocksTxsStorage {
	blocskStore := &BlocksTxsStorage{
		blocks: make([]*types.Block, 100),
		txs:    make(map[string]*types.Transaction),
		logger: logger,
	}
	// run subs
	go blocskStore.subscribeBlocks(cli)

	return blocskStore
}

// SubscribeBlocks subscribes blocks via infura API
func (b *BlocksTxsStorage) subscribeBlocks(cli *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := cli.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			b.logger.WithFields(logrus.Fields{"function": "SubscribeBlocks()"}).Fatalln(err)
			return
		case header := <-headers:
			block, err := cli.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				b.logger.WithFields(logrus.Fields{"function": "SubscribeBlocks()", "Method": "BlockByHash()"}).Errorln(err)
				time.Sleep(3 * time.Second)
				continue
			}

			b.Lock()
			// delete the first blocks(oldest) and it's transactions
			if len(b.blocks) == 100 {
				fmt.Println("HERE")
				oldBlock := b.blocks[0]
				b.blocks = b.blocks[1:]
				for _, tx := range oldBlock.Transactions() {
					delete(b.txs, tx.Hash().Hex())
				}
			}
			// append new block, it's transaction
			b.blocks = append(b.blocks, block)
			for _, tx := range block.Transactions() {
				b.txs[tx.Hash().Hex()] = tx
			}
			b.Unlock()

			b.logger.WithFields(logrus.Fields{
				"function":    "SubscribeBlocks()",
				"blockHash":   block.Hash().Hex(),
				"blockNumber": block.Number().Uint64(),
				"txs count":   len(block.Transactions()),
			}).Debugln("New Block")
		}
	}
}

// GetBlocksByTransactionID returns block contains transaction
// Input: hash of tranasction
// Output: block contains transaction with request hash. If success -> block, else -> nil
func (b *BlocksTxsStorage) GetBlockByTransactionID(id string) *types.Block {
	b.Lock()
	blocks := b.blocks
	b.Unlock()

	for _, bl := range blocks {
		for _, tx := range bl.Transactions() {
			if tx.Hash().Hex() == id {
				return bl
			}
		}
	}

	return nil
}

// GetBlocksByRange returns blocks from range
// Input: start block number, end block number of searching range
// Output: blocks and error. If success -> (blocks, nil), else -> (nil, error)
func (b *BlocksTxsStorage) GetBlocksByRange(start, end uint64) (rangeBlocks []*types.Block, err error) {
	b.Lock()
	blocks := b.blocks
	b.Unlock()

	if len(blocks) > 0 {
		firstBlock, lastBlock := blocks[0], blocks[len(blocks)-1]
		firstBlkNumber, lastBlkNumber := firstBlock.Number().Uint64(), lastBlock.Number().Uint64()
		if firstBlkNumber > start || lastBlkNumber < end {
			return nil, fmt.Errorf("Searching range (%d,%d) out of last 100 blocks range (%d,%d) ", start, end, firstBlkNumber, lastBlkNumber)
		}

	}

	return nil, fmt.Errorf("Please wait. Haven't got blocks yet")
}
