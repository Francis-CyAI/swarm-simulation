package main

import (
	"fmt"
	"swarm_server/node"
	"swarm_server/space"
	// "os"
)

func main() {
	/*
		// fmt.Println("\nWelcome to the CLI version of Swarm Simulation")
		var sp space.Plane
		sp_o := sp.Origin // (0, 0, 0) by default
		sp.End.X, sp.End.Y, sp.End.Z = 1, 1, 100

		var n node.Node

		// fmt.Printf("Node in main, before: %d, %d, %d\n", n.P.X, n.P.Y, n.P.Z)

		n.P = sp_o

		// fmt.Printf("Node in main, after: %d, %d, %d\n", n.P.X, n.P.Y, n.P.Z)

		occupiedSpace := []space.Point{sp_o}
		var done bool

		fmt.Printf("Current point is (%d, %d, %d)\n", n.P.X, n.P.Y, n.P.Z)

		if occupiedSpace, done = n.MoveTo(occupiedSpace, sp.End); done {
			fmt.Printf("Node in main, now at: %d, %d, %d\n", n.P.X, n.P.Y, n.P.Z)
			return
		}
	*/

	tester()
}

func tester() {
	var n node.Node
	limit := space.Point{X: 1, Y: 1, Z: 20}
	n.MoveTo([]space.Point{}, limit)
	fmt.Printf("n initial values; X: %d, Y: %d, Z: %d", n.P.X, n.P.Y, n.P.Z)

}
