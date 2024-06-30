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

	// free previously occupied point
	there, pIndex := node.IndexIn(space)
	if there {
		fmt.Printf("Node is there in space at index %d, now removing it\n", pIndex)
		space = slices.Delete(space, pIndex, pIndex+1)
		there, _ := node.IndexIn(space)
		fmt.Printf("There: %t\n", there)
	}

	var mu sync.Mutex

	var done bool
	mu.Lock()
	node.P, done = NextAvailableSpacePoint(space, node.P, limit)
	mu.Unlock()

	// fmt.Printf("pIndex = %d, pIndex + 1 = %d", pIndex, pIndex+1)
	// fmt.Printf("\tspace length = %d\n", len(space))

	return space, done
}

func NextAvailableSpacePoint(space []space.Point, p space.Point, limit space.Point) (space.Point, bool) {
	done := true

	fmt.Printf("node point in NextAvailableSpacePoint func, pre-op: %d %d %d\n", p.X, p.Y, p.Z)

	/*
	if p.X <= p.Y && p.X < limit.X || p.X <= p.Z && p.X < limit.X { // TODO: Fix this logic
		p.X += 1 
	} else if p.Y <= p.X && p.Y < limit.Y || p.Y <= p.Z  && p.Y < limit.Y {
		p.Y += 1
	} else if p.Z <= p.X && p.Z < limit.Z || p.Z <= p.Y && p.Z < limit.Z {
		p.Z += 1
	} else {
		return p, done
	} // limit reached
	*/

	if p.X < limit.X {
		if p.X <= p.Y && p.X <= p.Z {
			p.X += 1
		}
	} else if p.Y < limit.Y {
		if p.Y <= p.X && p.Y <= p.Z {
			p.Y += 1
		}
	} else if p.Z < limit.Z {
		if p.Z <= p.X && p.Z <= p.Y {
			p.Z += 1
		}
	} else {
		return p, done
	} // limit reached

	fmt.Printf("node point in NextAvailableSpacePoint func, post-op: %d %d %d\n\n", p.X, p.Y, p.Z)

	pointNotFree, _ := p.IndexIn(space)
	fmt.Printf("New point not free: %t\n", pointNotFree)
	if !pointNotFree { // ??

		// Point is not free
		NextAvailableSpacePoint(space, p, limit)
	}

	return p, !done // Not exactly useful, func just needs to return. Actual return statement at the top
}
