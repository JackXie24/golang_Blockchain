package main

import (
	"fmt"

	"github.com/JackXie24/golang_Blockchain/blockchain"
)

func main() {
	firstBlockChain := blockchain.InitBlockChain()
	firstBlockChain.AddBlock("This is the First Block of my first BlockChain")
	firstBlockChain.AddBlock("This is the Second Block of my first BlockChain, what purpose should my Blockchain serve")

	for _, block := range firstBlockChain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Perivous Hash: %x\n", block.Prev_Hash)
	}
}
