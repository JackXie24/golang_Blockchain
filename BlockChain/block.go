package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash      []byte
	Data      []byte
	Prev_Hash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Prev_Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, pre_hash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), pre_hash}
	block.DeriveHash()
	return block
}

func (blockChain *BlockChain) AddBlock(data string) {
	lastBlockHash := blockChain.Blocks[len(blockChain.Blocks)-1].Hash
	newBlock := CreateBlock(data, lastBlockHash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
