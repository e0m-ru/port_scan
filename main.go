package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"

	"e0m.ru/tcp_scaner/format"
	. "e0m.ru/tcp_scaner/loger"
	"github.com/cheggaaa/pb"
)

const ()

var (
	MAX_PORT          = int(math.Pow(2, 16)) - 1
	defaultPortString = fmt.Sprintf("1-%d", MAX_PORT)
	workersCount      int
	timeout           int
	portString        string
	portsRange        []int
	err               error
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
		flag.PrintDefaults()
	}
	flag.IntVar(&workersCount, "w", 100, "Determines the number of workers")
	flag.IntVar(&timeout, "t", 1000, "Determines the timeout for connection in miliseconds")
	flag.StringVar(&portString, "p", defaultPortString, "Ports define like -p [8080 || 1-1024 || 80,443,21,22]")
	flag.Parse()
	portsRange, err = format.Parse(portString)
	if err != nil {
		fmt.Printf("Can't parse port %v", portString)
		L.Fatalf("Can't parse port %v", portString)
	}
	if flag.Arg(0) == "" {
		fmt.Println("Not enough parameter\nusage: tcp_scaner [-wpt]... URL")
		os.Exit(0)
	}
}

func main() {
	address := fmt.Sprint(flag.Arg(0) + ":%v")
	fmt.Printf(address, portString)
	ports := make(chan int, workersCount)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}
	go func() {
		for _, i := range portsRange {
			ports <- i
		}
	}()

	bar := pb.StartNew(len(portsRange))

	for range portsRange {
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
