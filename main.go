package main

import (
	"fmt"

	"github.com/paolochang/gocoin/blockchain"
)

func main(){
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	chain.AddBlock("Forth block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}