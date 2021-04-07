package blockchain

import (
	"bytes"
	"crypto/sha256"
)

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

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}


