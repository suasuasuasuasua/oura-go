package main

import (
	"fmt"
	"net/http"
	"strings"
)

// WebHandler handles HTTP requests for the data visualization functionality
type WebHandler struct {
	chartGenerator *ChartGenerator
}

// NewWebHandler creates a new web handler
func NewWebHandler() *WebHandler {
	return &WebHandler{
		chartGenerator: NewChartGenerator(),
	}
}

// HomeHandler serves the main upload page
func (wh *WebHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Oura Data Visualization</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .upload-form { border: 2px dashed #ccc; padding: 20px; text-align: center; margin: 20px 0; }
        .chart-options { margin: 20px 0; }
        .form-group { margin: 10px 0; }
        label { display: inline-block; width: 120px; font-weight: bold; }
        select, input[type="text"] { width: 200px; padding: 5px; }
        button { padding: 10px 20px; background: #007cba; color: white; border: none; cursor: pointer; }
        button:hover { background: #005a87; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Oura Data Visualization</h1>
        <p>Upload your Oura CSV data file to create interactive charts and visualizations.</p>
        
        <form action="/upload" method="post" enctype="multipart/form-data" class="upload-form">
            <h3>Upload CSV File</h3>
            <input type="file" name="csvfile" accept=".csv" required>
            <br><br>
            
            <div class="chart-options">
                <h4>Chart Configuration</h4>
                <div class="form-group">
                    <label for="chartType">Chart Type:</label>
                    <select name="chartType" id="chartType">
                        <option value="line">Line Chart</option>
                        <option value="bar">Bar Chart</option>
                    </select>
                </div>
                
                <div class="form-group">
                    <label for="xColumn">X-Axis Column:</label>
                    <input type="text" name="xColumn" id="xColumn" placeholder="e.g., date" required>
                </div>
                
                <div class="form-group">
                    <label for="yColumn">Y-Axis Column:</label>
                    <input type="text" name="yColumn" id="yColumn" placeholder="e.g., sleep_score" required>
                </div>
                
                <div class="form-group">
                    <label for="title">Chart Title:</label>
                    <input type="text" name="title" id="title" placeholder="My Oura Data Chart">
                </div>
            </div>
            
            <button type="submit">Generate Chart</button>
        </form>
        
        <div style="margin-top: 40px;">
            <h3>Supported Data Columns</h3>
            <p>Common Oura CSV columns you can use:</p>
            <ul>
                <li><strong>date</strong> - Date of the measurement</li>
                <li><strong>sleep_score</strong> - Overall sleep score</li>
                <li><strong>total_sleep_duration</strong> - Total sleep time</li>
                <li><strong>efficiency</strong> - Sleep efficiency percentage</li>
                <li><strong>restfulness</strong> - Sleep restfulness score</li>
                <li><strong>rem_sleep_duration</strong> - REM sleep time</li>
                <li><strong>deep_sleep_duration</strong> - Deep sleep time</li>
                <li><strong>activity_score</strong> - Daily activity score</li>
                <li><strong>steps</strong> - Daily step count</li>
                <li><strong>readiness_score</strong> - Readiness score</li>
            </ul>
        </div>
    </div>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, tmpl)
}

// UploadHandler handles CSV file upload and chart generation
func (wh *WebHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get the file
	file, header, err := r.FormFile("csvfile")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".csv") {
		http.Error(w, "Please upload a CSV file", http.StatusBadRequest)
		return
	}

	// Get form parameters
	chartType := r.FormValue("chartType")
	xColumn := r.FormValue("xColumn")
	yColumn := r.FormValue("yColumn")
	title := r.FormValue("title")

	if title == "" {
		title = fmt.Sprintf("%s vs %s", yColumn, xColumn)
	}

	// Process CSV data
	csvData, err := ProcessCSV(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process CSV: %v", err), http.StatusBadRequest)
		return
	}

	// Generate chart based on type
	switch chartType {
	case "line":
		lineChart, err := wh.chartGenerator.CreateLineChart(csvData, xColumn, yColumn, title)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate chart: %v", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		if err := lineChart.Render(w); err != nil {
			http.Error(w, fmt.Sprintf("Failed to render chart: %v", err), http.StatusInternalServerError)
			return
		}
	case "bar":
		barChart, err := wh.chartGenerator.CreateBarChart(csvData, xColumn, yColumn, title)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate chart: %v", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		if err := barChart.Render(w); err != nil {
			http.Error(w, fmt.Sprintf("Failed to render chart: %v", err), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Invalid chart type", http.StatusBadRequest)
		return
	}
}

// InfoHandler provides information about available CSV columns
func (wh *WebHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	info := `
<!DOCTYPE html>
<html>
<head>
    <title>CSV Column Information</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        table { width: 100%; border-collapse: collapse; margin: 20px 0; }
        th, td { border: 1px solid #ddd; padding: 12px; text-align: left; }
        th { background-color: #f2f2f2; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Oura CSV Data Columns</h1>
        <p><a href="/">← Back to Upload</a></p>
        
        <h2>Common Oura Data Columns</h2>
        <table>
            <tr><th>Column Name</th><th>Description</th><th>Data Type</th></tr>
            <tr><td>date</td><td>Date of measurement</td><td>String (YYYY-MM-DD)</td></tr>
            <tr><td>sleep_score</td><td>Overall sleep quality score (0-100)</td><td>Numeric</td></tr>
            <tr><td>total_sleep_duration</td><td>Total sleep time in seconds</td><td>Numeric</td></tr>
            <tr><td>efficiency</td><td>Sleep efficiency percentage</td><td>Numeric</td></tr>
            <tr><td>restfulness</td><td>Sleep restfulness score</td><td>Numeric</td></tr>
            <tr><td>rem_sleep_duration</td><td>REM sleep duration in seconds</td><td>Numeric</td></tr>
            <tr><td>deep_sleep_duration</td><td>Deep sleep duration in seconds</td><td>Numeric</td></tr>
            <tr><td>light_sleep_duration</td><td>Light sleep duration in seconds</td><td>Numeric</td></tr>
            <tr><td>activity_score</td><td>Daily activity score (0-100)</td><td>Numeric</td></tr>
            <tr><td>steps</td><td>Daily step count</td><td>Numeric</td></tr>
            <tr><td>equivalent_walking_distance</td><td>Walking distance in meters</td><td>Numeric</td></tr>
            <tr><td>total_calories</td><td>Total calories burned</td><td>Numeric</td></tr>
            <tr><td>active_calories</td><td>Active calories burned</td><td>Numeric</td></tr>
            <tr><td>readiness_score</td><td>Readiness score (0-100)</td><td>Numeric</td></tr>
            <tr><td>temperature_deviation</td><td>Body temperature deviation</td><td>Numeric</td></tr>
            <tr><td>average_heart_rate</td><td>Average heart rate (BPM)</td><td>Numeric</td></tr>
            <tr><td>lowest_heart_rate</td><td>Lowest heart rate (BPM)</td><td>Numeric</td></tr>
        </table>
        
        <h2>Usage Tips</h2>
        <ul>
            <li>Use <strong>date</strong> as X-axis for time series data</li>
            <li>Numeric columns work best for Y-axis (scores, durations, counts)</li>
            <li>Column names are case-insensitive</li>
            <li>The system will skip rows with missing or invalid data</li>
        </ul>
    </div>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, info)
}

// RegisterRoutes registers all web handler routes
func (wh *WebHandler) RegisterRoutes() {
	http.HandleFunc("/", wh.HomeHandler)
	http.HandleFunc("/upload", wh.UploadHandler)
	http.HandleFunc("/info", wh.InfoHandler)
}