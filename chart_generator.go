package main

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// ChartGenerator handles chart creation using go-echarts
type ChartGenerator struct{}

// NewChartGenerator creates a new chart generator
func NewChartGenerator() *ChartGenerator {
	return &ChartGenerator{}
}

// CreateLineChart generates a line chart from CSV data
func (cg *ChartGenerator) CreateLineChart(csvData *CSVData, xColumn, yColumn, title string) (*charts.Line, error) {
	// Extract data
	xValues, err := csvData.GetStringColumn(xColumn)
	if err != nil {
		return nil, fmt.Errorf("failed to get X column data: %w", err)
	}
	
	yValues, err := csvData.GetNumericColumn(yColumn)
	if err != nil {
		return nil, fmt.Errorf("failed to get Y column data: %w", err)
	}
	
	// Ensure we have matching data lengths
	minLen := len(xValues)
	if len(yValues) < minLen {
		minLen = len(yValues)
	}
	
	if minLen == 0 {
		return nil, fmt.Errorf("no data available for chart")
	}
	
	// Create line chart
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xColumn,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yColumn,
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
		}),
	)
	
	// Prepare X-axis data
	xAxisData := make([]string, minLen)
	for i := 0; i < minLen; i++ {
		xAxisData[i] = xValues[i]
	}
	line.SetXAxis(xAxisData)
	
	// Prepare Y-axis data
	lineItems := make([]opts.LineData, minLen)
	for i := 0; i < minLen; i++ {
		lineItems[i] = opts.LineData{Value: yValues[i]}
	}
	
	line.AddSeries(yColumn, lineItems)
	
	return line, nil
}

// CreateBarChart generates a bar chart from CSV data
func (cg *ChartGenerator) CreateBarChart(csvData *CSVData, xColumn, yColumn, title string) (*charts.Bar, error) {
	// Extract data
	xValues, err := csvData.GetStringColumn(xColumn)
	if err != nil {
		return nil, fmt.Errorf("failed to get X column data: %w", err)
	}
	
	yValues, err := csvData.GetNumericColumn(yColumn)
	if err != nil {
		return nil, fmt.Errorf("failed to get Y column data: %w", err)
	}
	
	// Ensure we have matching data lengths
	minLen := len(xValues)
	if len(yValues) < minLen {
		minLen = len(yValues)
	}
	
	if minLen == 0 {
		return nil, fmt.Errorf("no data available for chart")
	}
	
	// Create bar chart
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xColumn,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yColumn,
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
		}),
	)
	
	// Prepare X-axis data
	xAxisData := make([]string, minLen)
	for i := 0; i < minLen; i++ {
		xAxisData[i] = xValues[i]
	}
	bar.SetXAxis(xAxisData)
	
	// Prepare Y-axis data
	barItems := make([]opts.BarData, minLen)
	for i := 0; i < minLen; i++ {
		barItems[i] = opts.BarData{Value: yValues[i]}
	}
	
	bar.AddSeries(yColumn, barItems)
	
	return bar, nil
}