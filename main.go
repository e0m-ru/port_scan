package main

import (
	"fmt"
	"math"
	"net"
	"sort"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.1.2:%d", p)
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
	n := int(math.Pow(2, 16))
	ports := make(chan int, 10000)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= n; i++ {
			ports <- i
		}
	}()

	bar := pb.StartNew(n)

	for i := 0; i < n; i++ {
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
