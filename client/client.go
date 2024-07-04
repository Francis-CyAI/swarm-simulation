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
	
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}