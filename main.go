package main

import (
	"github.com/paolochang/gocoin/blockchain"
)

func main() {
	// cli.Strat()
	blockchain.Blockchain().AddBlock("first")
	blockchain.Blockchain().AddBlock("second")
	blockchain.Blockchain().AddBlock("third")
}

