package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Load the configuration file
	config := loadConfig("config.json")

	// Initialize the blockchain network
	network, err := initializeNetwork(config.Nodes)
	if err != nil {
		fmt.Println("Error initializing network:", err)
		os.Exit(1)
	}

	// Run the load tests
	startTime := time.Now()
	testResults := runLoadTests(config, network)
	endTime := time.Now()

	// Generate the test report
	report := generateReport(config, testResults, startTime, endTime)

	// Print the report to console and save to file
	fmt.Println(report)
	saveReportToFile(report, "report.txt")
}
