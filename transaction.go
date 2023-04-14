package main

import (
	"errors"
)

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	Sender   string `json:"sender"`    // sender of the transaction
	Receiver string `json:"receiver"`  // receiver of the transaction
	Amount   int    `json:"amount"`    // amount of the transaction
}

// TransactionPool represents a pool of pending transactions.
type TransactionPool struct {
	Pool []Transaction `json:"pool"`  // list of pending transactions
}

// NewTransaction creates a new transaction.
func NewTransaction(sender string, receiver string, amount int) Transaction {
	return Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

// AddTransaction adds a new transaction to the transaction pool.
func (tp *TransactionPool) AddTransaction(tx Transaction) {
	tp.Pool = append(tp.Pool, tx)
}

// GetTransaction returns a transaction from the transaction pool.
func (tp *TransactionPool) GetTransaction(sender string, receiver string, amount int) (Transaction, error) {
	for _, tx := range tp.Pool {
		if tx.Sender == sender && tx.Receiver == receiver && tx.Amount == amount {
			return tx, nil
		}
	}
	return Transaction{}, errors.New("transaction not found")
}

// RemoveTransaction removes a transaction from the transaction pool.
func (tp *TransactionPool) RemoveTransaction(tx Transaction) {
	for i, t := range tp.Pool {
		if t == tx {
			tp.Pool = append(tp.Pool[:i], tp.Pool[i+1:]...)
			return
		}
	}
}

// ClearPool clears the transaction pool.
func (tp *TransactionPool) ClearPool() {
	tp.Pool = nil
}
