package blockchain

import (
	"sync"

	"github.com/paolochang/gocoin/db"
	"github.com/paolochang/gocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
}

// Singleton Pattern
var b * blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	// block := Block{data, "", b.NewestHash, b.Height + 1}
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Initial_block")
		})
	}
	return b
}
