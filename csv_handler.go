package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// CSVData represents processed CSV data
type CSVData struct {
	Headers []string
	Rows    [][]string
}

// ProcessCSV reads and processes CSV data from a reader
func ProcessCSV(reader io.Reader) (*CSVData, error) {
	csvReader := csv.NewReader(reader)
	
	// Read all records
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}
	
	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}
	
	return &CSVData{
		Headers: records[0],
		Rows:    records[1:],
	}, nil
}

// GetNumericColumn extracts numeric values from a specific column
func (c *CSVData) GetNumericColumn(columnName string) ([]float64, error) {
	// Find column index
	columnIndex := -1
	for i, header := range c.Headers {
		if strings.EqualFold(header, columnName) {
			columnIndex = i
			break
		}
	}
	
	if columnIndex == -1 {
		return nil, fmt.Errorf("column '%s' not found", columnName)
	}
	
	var values []float64
	for _, row := range c.Rows {
		if columnIndex >= len(row) {
			continue
		}
		
		value, err := strconv.ParseFloat(strings.TrimSpace(row[columnIndex]), 64)
		if err != nil {
			// Skip non-numeric values
			continue
		}
		values = append(values, value)
	}
	
	return values, nil
}

// GetStringColumn extracts string values from a specific column
func (c *CSVData) GetStringColumn(columnName string) ([]string, error) {
	// Find column index
	columnIndex := -1
	for i, header := range c.Headers {
		if strings.EqualFold(header, columnName) {
			columnIndex = i
			break
		}
	}
	
	if columnIndex == -1 {
		return nil, fmt.Errorf("column '%s' not found", columnName)
	}
	
	var values []string
	for _, row := range c.Rows {
		if columnIndex >= len(row) {
			continue
		}
		values = append(values, strings.TrimSpace(row[columnIndex]))
	}
	
	return values, nil
}