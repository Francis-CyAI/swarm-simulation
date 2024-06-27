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

func (node *Node) MoveTo(space []space.Point, p space.Point){
	/**
		Node must first check if point exists in
		the occupied points Slice. If it does, it
		chooses a different point.
	*/
	var mu sync.Mutex
	
	mu.Lock()	
	if pointNotFree := slices.Contains(space, p); !pointNotFree {
		node.P = p
	}
	defer mu.Unlock()

}