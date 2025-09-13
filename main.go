package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/netip"
)

var verbose bool
var addrPort netip.AddrPort

func init() {
	// see https://pkg.go.dev/flag
	const (
		defaultDebug   = false
		defaultAddress = "127.0.0.1"
		defaultPort    = 8080
	)

	flag.BoolVar(&verbose, "v", defaultDebug, "enable verbose mode (shorthand)")
	flag.BoolVar(&verbose, "verbose", defaultDebug, "enable verbose mode")

	var addr = flag.String("address", defaultAddress, "the ip address to server the webserver on")
	var port = flag.Uint("port", defaultPort, "the port to serve webserver on")

	// combine the port and address into the netip.AddrPort struct
	var address, err = netip.ParseAddr(*addr)
	if err != nil {
		log.Fatalf("Address '%v' is not vaild\n", addr)
	}
	addrPort = netip.AddrPortFrom(address, uint16(*port))
}

func main() {
	// Parse the flags
	flag.Parse()

	if verbose {
		fmt.Printf("Starting Oura data visualization server on port %d\n", addrPort.Port())
		fmt.Printf("Visit http://%s to upload CSV files and create charts\n", addrPort.String())
	}

	// Initialize web handler for data visualization
	webHandler := NewWebHandler()
	webHandler.RegisterRoutes()

	fmt.Println("Oura Go Data Visualization Server")
	fmt.Println("==================================")
	fmt.Printf("Server running on: http://%s\n", addrPort.String())
	fmt.Println("Features:")
	fmt.Println("  - CSV file upload")
	fmt.Println("  - Interactive charts (line/bar)")
	fmt.Println("  - Oura health data visualization")
	fmt.Println()

	var err = http.ListenAndServe(addrPort.String(), nil)
	// NOTE: even though http.ListenAndServe returns an error, the docs
	// says that it will _always_ return a non-nill error
	// this is for completeness so my linter is happy
	if err != nil {
		log.Fatalln("Something horrible has gone wrong!")
	}
}
