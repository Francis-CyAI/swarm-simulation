package main

import (
	"fmt"
	"swarm_server/node"
	"swarm_server/space"
	// "os"
)

func main() {
		fmt.Println("\nWelcome to the CLI version of Swarm Simulation")

		var (
			sp = space.Plane{End: space.Point{X: 1, Y: 1000, Z: 3}}
			occupiedSpace []space.Point
		)

		numOfNodes := 3

		ch := make(chan struct{}, numOfNodes) // Needed to make main go routine wait for subroutines

		for range(numOfNodes) {
			go func() {
				var n node.Node
				n.MoveTo(occupiedSpace, sp.End)
				ch<-struct{}{}
			}()
		}

		for range(numOfNodes) {
			<-ch
		}

	// tester()
}

func tester() {
	ch := make(chan struct{})
	var n node.Node
	limit := space.Point{X: 1, Z: 1, Y: 20}
	go func() {
		n.MoveTo([]space.Point{}, limit)
		ch <- struct{}{}
	}()
	fmt.Printf("n initial values; X: %d, Y: %d, Z: %d", n.P.X, n.P.Y, n.P.Z)
	<-ch

}
