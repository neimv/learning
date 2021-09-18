package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()
	mapas := make(map[int]string[])
	fmt.Println("vim-go")
	var wg sync.WaitGroup

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			// Check ports in this range
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))

			if err != nil {
				// fmt.Printf("Is not possible connect to %d\n", i)
				return
			}

			conn.Close()
			fmt.Printf("The port %d is open\n", port)
		}(i)
	}
	wg.Wait()
}
