package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBlockchainLoadTesting(t *testing.T) {
	config := NewConfig(100, 5, 2, 1, "localhost:8080")
	blockchain := NewBlockchain(config)

	// Start monitoring the blockchain network
	monitor := NewMonitor(blockchain, time.Second*5)
	go monitor.Start()
	defer monitor.Stop()

	// Generate transactions and add to the blockchain
	for i := 1; i <= config.Transactions; i++ {
		transaction := NewTransaction(fmt.Sprintf("Sender%d", i), fmt.Sprintf("Recipient%d", i), 1)
		blockchain.AddTransaction(transaction)
	}

	// Start mining blocks
	blockchain.MineBlocks()

	// Generate report
	report := NewReport("performance_report.csv", blockchain.GetPerformanceData())
	err := report.GenerateCSV()
	if err != nil {
		t.Errorf("Error generating report: %v", err)
	}
}
