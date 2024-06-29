package node

import (
	"swarm_server/space"
	"sync"
	"slices"
)

type Node struct {
	R float32 /// Radius
	P space.Point /// Position of center
}

func (node *Node) MoveTo(space []space.Point) []space.Point {
	/**
		Node must first check if point exists in
		the occupied points Slice. If it does, it
		chooses a different point.
	*/
	var mu sync.Mutex
	
	mu.Lock()	
	node.P = NextAvailableSpacePoint(space, node.P)
	mu.Unlock()
	
	// remove previous point
	pIndex := slices.Index(space, node.P)
	space = slices.Delete(space, pIndex, pIndex + 1)

	return space
}

func NextAvailableSpacePoint(space []space.Point, p space.Point) space.Point {
	if p.X <= p.Y || p.X <= p.Z {
		p.X += 1
	} else if p.Y <= p.X || p.Y <= p.Z {
		p.Y += 1
	} else {
		p.Z += 1
	}

	if pointNotFree := slices.Contains(space, p); pointNotFree {
		// Point is not free
		NextAvailableSpacePoint(space, p)
	}

	return p
}