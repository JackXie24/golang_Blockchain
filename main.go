package main

import (
	"fmt"
)

func main() {
	firstBlockChain := blockChain.initBlockChain()
	firstBlockChain.addBlock("This is the First Block of my first BlockChain")
	firstBlockChain.addBlock("This is the Second Block of my first BlockChain, what purpose should my Blockchain serve")

	for _, block := range firstBlockChain.blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Perivous Hash: %x\n", block.Prev_Hash)
	}
}
