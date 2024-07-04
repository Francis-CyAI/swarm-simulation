package main

import (
	"fmt"
	"log"
	"net"
	"swarm_server/node"
	"swarm_server/space"
	"sync"
)

func main() {
	fmt.Println("\nSwarm server running...")

	var (
		sp            = space.Plane{End: space.Point{X: 1, Y: 500000, Z: 3}}
		occupiedSpace []space.Point
	)

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	var (
		wg sync.WaitGroup

		mu    sync.Mutex
		count int
	)

	for {
		conn, err := listener.Accept()
		// conn closed in MoveTo method
		if err != nil {
			log.Print(err)
			break
		}

		wg.Add(1)

		go func() {
			mu.Lock()
			count++
			num := count
			mu.Unlock()
			defer wg.Done()
			var n node.Node
			point, done := n.MoveTo(conn, occupiedSpace, sp.End)
			fmt.Printf("Node # %d: done: %t, at: (%d, %d, %d)\n", num, done, point.X, point.Y, point.Z)
			count--
		}()
	}
	wg.Wait()
	fmt.Println("\nShutdown.")
}
