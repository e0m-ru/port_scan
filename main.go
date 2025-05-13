package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	MAX_PORT = 65535
)

var (
	defaultPortString = fmt.Sprintf("1-%d", MAX_PORT)
	workersCount      int
	timeout           int
	portString        string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
		flag.PrintDefaults()
	}
	flag.IntVar(&workersCount, "w", 100, "Determines the number of workers")
	flag.IntVar(&timeout, "t", 1000, "Determines the timeout for connection in milliseconds")
	flag.StringVar(&portString, "p", defaultPortString, "Ports define like -p [8080 || 1-1024 || 1-80,443,21-22,4455]")
}

func main() {
	// Parse flags
	flag.Parse()

	// Check arguments
	if flag.Arg(0) == "" {
		fmt.Println("Error: Missing target URL\nUsage: tcp_scaner [-wpt]... URL")
		os.Exit(1)
	}
	target := flag.Arg(0)

	// Parse port range
	portsRange, err := ParsePortRanges(portString)
	if err != nil {
		log.Fatalf("Error parsing ports: %v", err)
	}

	var (
		wg sync.WaitGroup
	)

	// Send ports to channel
	ports := make(chan int, len(portsRange))
	go func() {
		for _, port := range portsRange {
			ports <- port
		}
		close(ports)
	}()

	// Start workers
	results := make(chan int)
	if len(portsRange) < workersCount {
		workersCount = len(portsRange)
	}
	for range workersCount {
		wg.Add(1)
		go worker(&wg, ports, results, target+":%d")
	}

	// Read results
	go func() {
		for port := range results {
			if port != 0 {
				fmt.Printf("%d ", port)
			}
		}
	}()
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup, ports, results chan int, address string) {
	d := net.Dialer{Timeout: time.Millisecond * time.Duration(timeout)}
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf(address, p)
		conn, err := d.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
