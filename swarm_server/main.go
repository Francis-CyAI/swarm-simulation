package main

import (
	"fmt"
	"swarm_server/node"
	"swarm_server/space"
	"sync"
)

func main() {
		fmt.Println("\nWelcome to the CLI version of Swarm Simulation")

		var (
			sp = space.Plane{End: space.Point{X: 1, Y: 1000, Z: 3}}
			occupiedSpace []space.Point
		)

		numOfNodes := 3
		
		var wg sync.WaitGroup

		for range(numOfNodes) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var n node.Node
				n.MoveTo(occupiedSpace, sp.End)
			}()
		}

		wg.Wait()
}
