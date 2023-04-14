package main

import (
	"fmt"
	"time"
)

// Monitor represents a monitor for measuring the performance of blockchain networks.
type Monitor struct {
	Blockchain *Blockchain // blockchain instance to monitor
	Interval   time.Duration // monitoring interval
	stop       chan bool // channel to signal stop monitoring
}

// NewMonitor creates a new monitor for measuring the performance of blockchain networks.
func NewMonitor(blockchain *Blockchain, interval time.Duration) *Monitor {
	return &Monitor{
		Blockchain: blockchain,
		Interval:   interval,
		stop:       make(chan bool),
	}
}

// Start starts monitoring the blockchain network.
func (m *Monitor) Start() {
	ticker := time.NewTicker(m.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get and print blockchain statistics
			stats := m.Blockchain.GetStatistics()
			fmt.Printf("Block height: %d | Transactions: %d | Mining Difficulty: %d | Hashrate: %d H/s\n",
				stats.BlockHeight, stats.Transactions, stats.MiningDifficulty, stats.HashRate)
		case <-m.stop:
			return
		}
	}
}

// Stop stops monitoring the blockchain network.
func (m *Monitor) Stop() {
	m.stop <- true
}
