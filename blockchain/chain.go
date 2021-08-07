package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
}

// Singleton Pattern
var b * blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	// block := Block{data, "", b.NewestHash, b.Height + 1}
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Initial block")
		})
	}
	return b
}
