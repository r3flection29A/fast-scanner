package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("\tPortscanner in Golang")
	if len(os.Args) < 2 {
		fmt.Println("Missing params...")
		fmt.Println("Ex: ./fast-scanner scanme.nmap.org")
		os.Exit(1)
	} else {
		fmt.Println("Target ------>", os.Args[1])
	}

	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addr := os.Args[1] + ":" + strconv.Itoa(j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			defer conn.Close()
			fmt.Printf("Port %d open\n", j)
		}(i)
	}
	wg.Wait()
}
