package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/netip"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {
	var _, err = fmt.Fprintf(w, "Hello")
	if err != nil {
		log.Fatalln("Could not write to http.ResponseWriter")
	}
}

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
		fmt.Printf("Starting server on port %d\n", addrPort.Port())
		fmt.Printf("Visit on http://%s:%d\n", addrPort.Addr(), addrPort.Port())
	}

	http.HandleFunc("/", httpserver)
	var err = http.ListenAndServe(addrPort.String(), nil)
	if err != nil {
		log.Fatalln("Something horrible has gone wrong!")
	}
}
