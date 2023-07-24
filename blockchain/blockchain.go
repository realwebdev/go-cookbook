package main

import (
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PreviousHash []byte
}

var blocks []Block

func NewBlock(data string, previousHash []byte) Block {
	block := Block{
		Index:        len(blocks),
		Timestamp:    time.Now(),
		Data:         data,
		PreviousHash: previousHash,
	}
	return block
}

func (block *Block) GetPreviousHash() []byte {
	return block.PreviousHash
}

func main() {
	blocks = []Block{} // Move blocks definition to here
	genesisBlock := NewBlock("Genesis Block", []byte{})
	blocks = append(blocks, genesisBlock)

	for i := 1; i < 10; i++ {
		block := NewBlock(fmt.Sprintf("Block %d", i), blocks[i-1].GetPreviousHash())
		blocks = append(blocks, block)
	}

	for _, block := range blocks {
		fmt.Println(block)
	}
}
