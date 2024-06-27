package main

import (
	"swarm_server/space"
	// "swarm_server/node"
	"fmt"
	// "os"
)

func main() {
	// fmt.Println("\nWelcome to the CLI version of Swarm Simulation")
	var sp space.Plane
	sp_o := sp.Origin // (0, 0, 0) by default
	
	fmt.Print(sp_o.X)
}
