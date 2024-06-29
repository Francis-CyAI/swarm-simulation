package node

import (
	"fmt" // TODO: remove
	"slices"
	"swarm_server/space"
	"sync"
)

type Node struct {
	R float32     /// Radius
	P space.Point /// Position of center
}

func (node Node) IndexIn(space []space.Point) (bool, int) {
	there := false
	i := -1

	for j := 0; j < len(space); j++ {
		/*
		fmt.Printf("space point = ")
		fmt.Printf(space[j])
		fmt.Printf("\nnode point = ")
		fmt.Printf(node.P)
		*/
		fmt.Printf("space[j] (space[%d]): X = %d, Y = %d, Z = %d\n", j, space[j].X, space[j].Y, space[j].Z)
		fmt.Printf("node: X = %d, Y = %d, Z = %d\n \n", node.P.X, node.P.Y, node.P.Z)

		if currSPoint := space[j]; node.P.X == currSPoint.X && node.P.Y == currSPoint.Y && node.P.Z == currSPoint.Z {
			there = true
			i = j
		}
	}

	return there, i
}

func (node Node) MoveTo(space []space.Point, limit space.Point) ([]space.Point, bool) { // TODO: fix logic
	/**
	Node must first check if point exists in
	the occupied points Slice. If it does, it
	chooses a different point.
	*/

	var mu sync.Mutex
	var done bool
	mu.Lock()
	node.P, done = NextAvailableSpacePoint(space, node.P, limit)
	mu.Unlock()

	// remove previous point
	there, pIndex := node.IndexIn(space)
	if there {
		space = slices.Delete(space, pIndex, pIndex+1)
	}
	// fmt.Printf("pIndex = %d, pIndex + 1 = %d", pIndex, pIndex+1)
	// fmt.Printf("\tspace length = %d\n", len(space))

	return space, done
}

func NextAvailableSpacePoint(space []space.Point, p space.Point, limit space.Point) (space.Point, bool) {
	done := true

	if p.X >= limit.X && p.Y >= limit.Y && p.Z >= limit.Z {
		return p, done
	} // limit reached

	if p.X <= p.Y || p.X <= p.Z {
		p.X += 1
	} else if p.Y <= p.X || p.Y <= p.Z {
		p.Y += 1
	} else {
		p.Z += 1
	}

	if pointNotFree := slices.Contains(space, p); pointNotFree {
		// Point is not free
		// NextAvailableSpacePoint(space, p, limit)
	}

	return p, !done
}
