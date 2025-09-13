package main

import (
	"strings"
	"testing"
)

func TestProcessCSV(t *testing.T) {
	csvData := `date,sleep_score,steps
2024-01-01,85,8500
2024-01-02,78,7200
2024-01-03,92,9100`

	reader := strings.NewReader(csvData)
	data, err := ProcessCSV(reader)
	if err != nil {
		t.Fatalf("ProcessCSV failed: %v", err)
	}

	// Check headers
	expectedHeaders := []string{"date", "sleep_score", "steps"}
	if len(data.Headers) != len(expectedHeaders) {
		t.Fatalf("Expected %d headers, got %d", len(expectedHeaders), len(data.Headers))
	}

	for i, header := range expectedHeaders {
		if data.Headers[i] != header {
			t.Errorf("Expected header %s, got %s", header, data.Headers[i])
		}
	}

	// Check rows
	if len(data.Rows) != 3 {
		t.Fatalf("Expected 3 rows, got %d", len(data.Rows))
	}

	// Test numeric column extraction
	scores, err := data.GetNumericColumn("sleep_score")
	if err != nil {
		t.Fatalf("GetNumericColumn failed: %v", err)
	}

	expectedScores := []float64{85, 78, 92}
	if len(scores) != len(expectedScores) {
		t.Fatalf("Expected %d scores, got %d", len(expectedScores), len(scores))
	}

	for i, score := range expectedScores {
		if scores[i] != score {
			t.Errorf("Expected score %f, got %f", score, scores[i])
		}
	}

	// Test string column extraction
	dates, err := data.GetStringColumn("date")
	if err != nil {
		t.Fatalf("GetStringColumn failed: %v", err)
	}

	expectedDates := []string{"2024-01-01", "2024-01-02", "2024-01-03"}
	if len(dates) != len(expectedDates) {
		t.Fatalf("Expected %d dates, got %d", len(expectedDates), len(dates))
	}

	for i, date := range expectedDates {
		if dates[i] != date {
			t.Errorf("Expected date %s, got %s", date, dates[i])
		}
	}
}

func TestChartGenerator(t *testing.T) {
	csvData := `date,sleep_score
2024-01-01,85
2024-01-02,78
2024-01-03,92`

	reader := strings.NewReader(csvData)
	data, err := ProcessCSV(reader)
	if err != nil {
		t.Fatalf("ProcessCSV failed: %v", err)
	}

	generator := NewChartGenerator()

	// Test line chart creation
	lineChart, err := generator.CreateLineChart(data, "date", "sleep_score", "Test Line Chart")
	if err != nil {
		t.Fatalf("CreateLineChart failed: %v", err)
	}
	if lineChart == nil {
		t.Fatal("CreateLineChart returned nil chart")
	}

	// Test bar chart creation
	barChart, err := generator.CreateBarChart(data, "date", "sleep_score", "Test Bar Chart")
	if err != nil {
		t.Fatalf("CreateBarChart failed: %v", err)
	}
	if barChart == nil {
		t.Fatal("CreateBarChart returned nil chart")
	}
}