package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	newBlock := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	newBlock.Hash = CalculateHash(transaction + previousHash)
	return newBlock
}

func (b *Blockchain) AddBlock(newBlock *Block) {
	if len(b.Blocks) > 0 {
		newBlock.PreviousHash = b.Blocks[len(b.Blocks)-1].Hash
	}
	b.Blocks = append(b.Blocks, newBlock)
}

func (b *Blockchain) DisplayBlocks() {
	for _, block := range b.Blocks {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}

func (b *Blockchain) ChangeBlock(index int, transaction string) {
	b.Blocks[index].Transaction = transaction
	b.Blocks[index].Hash = CalculateHash(transaction + b.Blocks[index].PreviousHash)
}

func (b *Blockchain) VerifyChain() bool {
	for i := 1; i < len(b.Blocks); i++ {
		if b.Blocks[i].Hash != CalculateHash(b.Blocks[i].Transaction+b.Blocks[i-1].Hash) || (i > 1 && b.Blocks[i].PreviousHash != b.Blocks[i-1].Hash) {
			return false
		}
	}
	return true
}

func CalculateHash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hashed := h.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
