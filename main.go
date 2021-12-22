package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var common = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}

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
