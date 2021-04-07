package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

const (
	dbpath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte
	opts := badger.DefultOptions
	opts.Dir = dbpath
	opts.ValueDir = dbpath

	db, err := badger.Open(opts)
	Handle(err)
	err := db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No exsiting blockchainFound")
			genisis := Genesis()
			fmt.Println("Genesis Proved")
			err = txn.Set(genisis.Hash, genisis.Serialize())
			lastHash = genisis.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)
	blockchain := BlockChain{lastHash, db}

	return &blockchain
}

func (blockChain *BlockChain) AddBlock(data string) {
	lastBlockHash := blockChain.Blocks[len(blockChain.Blocks)-1].Hash
	newBlock := CreateBlock(data, lastBlockHash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
