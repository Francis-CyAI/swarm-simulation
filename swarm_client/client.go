package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Simply connects and log all responses form the
	// server to standard output

	/*
		num, err := strconv.ParseUint(os.Args[1:2][0], 10, 64)
		if err != nil {
			fmt.Println("An error has occured. Please start client with a different type of number")
			return
		}
	*/

	// var wg sync.WaitGroup
	/*
		for i := 1; i <= int(num); i++ {
			fmt.Printf("client # %d", i)
			wg.Add(1)
			go func() {
	*/
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// defer wg.Done()

	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
	// }()
	// }
	// wg.Wait()

	// swarmClient()
}
