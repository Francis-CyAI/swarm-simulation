package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Simply connects and log all responses form the
	// server to standard output

	/*
		Spawning multiple goroutines from a single process causes them to be multiplexed, 
		so indidual processes must be started from separate terminal instances.
	*/


	conn := estConn(bufio.NewReader(os.Stdin))
	defer conn.Close()

	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}

func estConn(reader *bufio.Reader) net.Conn {
	fmt.Println("\nPlease enter the network address for the swarm server, e.g., localhost:8000, and press Enter to connect")


	netAddr, _ := reader.ReadString('\n')

	conn, err := net.Dial("tcp", strings.TrimSpace(netAddr))
	if err != nil {
		log.Fatal(err)
		estConn(reader)
	}
	return conn
}