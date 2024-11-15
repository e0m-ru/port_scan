package main

import (
	"fmt"
	"net"
	"time"
)

func worker(ports, results chan int, address string) {
	for p := range ports {
		address := fmt.Sprintf(address, p)
		d := net.Dialer{Timeout: time.Millisecond * time.Duration(timeout)}
		conn, err := d.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
