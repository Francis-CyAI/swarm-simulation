package main

import (
	"fmt"
	"swarm_server/node"
	"swarm_server/space"
	// "os"
)

func main() {
	/*
	fmt.Println("\nWelcome to the CLI version of Swarm Simulation")

	var (
		sp = space.Plane{End: space.Point{X: 1, Y: 3, Z: 3}}
		
		occupiedSpace []space.Point

		n1 node.Node
		// n2 node.Node
		// n3 node.Node
	)

	go n1.MoveTo(occupiedSpace, sp.End)
	// go n2.MoveTo(occupiedSpace, sp.End)
	// go n3.MoveTo(occupiedSpace, sp.End)
*/

	tester()
}

func tester() {
	var n node.Node
	limit := space.Point{X: 1, Z: 1, Y: 20}
	n.MoveTo([]space.Point{}, limit)
	fmt.Printf("n initial values; X: %d, Y: %d, Z: %d", n.P.X, n.P.Y, n.P.Z)

}
