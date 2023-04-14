package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config represents the configuration parameters for the load testing framework.
type Config struct {
	Nodes           int     `json:"nodes"`             // number of nodes in the blockchain network
	Transactions    int     `json:"transactions"`      // number of transactions to generate
	TPS             int     `json:"tps"`               // target transactions per second
	NetworkLatency  int     `json:"network_latency"`   // network latency in milliseconds
	MaxBlocksize    int     `json:"max_blocksize"`     // maximum block size in bytes
	MaxTxPerBlock   int     `json:"max_tx_per_block"`  // maximum transactions per block
	ResourceMonitor bool    `json:"resource_monitor"`  // whether to monitor resource usage during testing
	OutputFile      string  `json:"output_file"`       // name of the output file for the load testing report
}

// loadConfig loads the configuration file and returns a Config struct.
func loadConfig(filename string) Config {
	var config Config

	// Read the config file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}

	// Parse the JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error parsing configuration data:", err)
	}

	return config
}
