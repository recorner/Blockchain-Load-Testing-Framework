package main

import (
	"errors"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Index        int           `json:"index"`         // index of the block in the blockchain
	Timestamp    int64         `json:"timestamp"`     // timestamp of the block
	Transactions []Transaction `json:"transactions"`  // list of transactions in the block
	PrevHash     string        `json:"prev_hash"`     // hash of the previous block in the chain
	Hash         string        `json:"hash"`          // hash of the current block
}

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	Sender   string `json:"sender"`    // sender of the transaction
	Receiver string `json:"receiver"`  // receiver of the transaction
	Amount   int    `json:"amount"`    // amount of the transaction
}

// Blockchain represents the blockchain.
type Blockchain struct {
	Chain        []Block `json:"chain"`         // list of blocks in the blockchain
	CurrentIndex int     `json:"current_index"` // index of the current block
}

// NewBlock creates a new block in the blockchain.
func NewBlock(index int, transactions []Transaction, prevHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(prevBlock.Index+1, transactions, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

// IsValid checks if the blockchain is valid.
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		prevBlock := bc.Chain[i-1]
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

// GetBalance returns the balance of a given account.
func (bc *Blockchain) GetBalance(account string) (int, error) {
	balance := 0
	for _, block := range bc.Chain {
		for _, tx := range block.Transactions {
			if tx.Sender == account {
				balance -= tx.Amount
			}
			if tx.Receiver == account {
				balance += tx.Amount
			}
		}
	}
	if balance < 0 {
		return 0, errors.New("negative balance")
	}
	return balance, nil
}

// calculateHash calculates the SHA-256 hash of a block.
func calculateHash(block Block) string {
	// implementation of SHA-256 hashing algorithm goes here
	return ""
}
