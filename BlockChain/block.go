package BlockChain

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

func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Prev_Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, pre_hash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), pre_hash}
	block.deriveHash()
	return block
}

func (blockChain *BlockChain) addBlock(data string) {
	lastBlockHash := blockChain.Blocks[len(blockChain.Blocks)-1].Hash
	newBlock := createBlock(data, lastBlockHash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func initBlockChain() *BlockChain {
	return &BlockChain{[]*Block{genesis()}}
}
