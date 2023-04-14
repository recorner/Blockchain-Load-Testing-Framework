package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Report represents a report for measuring the performance of blockchain networks.
type Report struct {
	FileName string // name of the report file
	Data     [][]string // data to write to the report file
}

// NewReport creates a new report for measuring the performance of blockchain networks.
func NewReport(fileName string, data [][]string) *Report {
	return &Report{
		FileName: fileName,
		Data:     data,
	}
}

// GenerateCSV generates a CSV file with the report data.
func (r *Report) GenerateCSV() error {
	file, err := os.Create(r.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range r.Data {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Report generated successfully: %s\n", r.FileName)
	return nil
}
