package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"sort"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	workersCount, portsCount int
)

func init() {
	flag.IntVar(&workersCount, "w", 10000, "Determines the number of workers")
	flag.IntVar(&portsCount, "p", int(math.Pow(2, 16)), `Defines the port range from 0 to N`)
}

func worker(ports, results chan int, address string) {
	for p := range ports {
		address := fmt.Sprintf(address, p)
		d := net.Dialer{Timeout: time.Second}
		conn, err := d.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
func main() {

	flag.Parse()
	address := fmt.Sprint(flag.Arg(0) + ":%d")
	fmt.Printf("%v\n", address)
	ports := make(chan int, workersCount)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}
	go func() {
		for i := 1; i <= portsCount; i++ {
			ports <- i
		}
	}()

	bar := pb.StartNew(portsCount)

	for i := 0; i < portsCount; i++ {
		bar.Increment()
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	bar.Finish()
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
