package node

import (
	// TODO: remove
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

/*
func indexIn(space []space.Point, p space.Point) (bool, int) {
	there := false
	i := -1
	for j := 0; j < len(space); j++ {
		if currSPoint := space[j]; p.X == currSPoint.X && p.Y == currSPoint.Y && p.Z == currSPoint.Z {
			there = true
			i = j
		}
	}

	return there, i
}
*/

func nextPoint(currPoint space.Point, limit space.Point) (space.Point, bool) {
	// return next point, and done status

	var done bool
	
	if limit.X >= limit.Y && limit.X >= limit.Z {
		// X is the main dimension
		currPoint.X++
		done = currPoint.X >= limit.X
	} else if limit.Y >= limit.X && limit.Y >= limit.Z {
		// Y is the main dimention
		currPoint.Y++
		done = currPoint.Y >= limit.Y
	} else if limit.Z >= limit.X && limit.Z >= limit.Y {
		// Z is the main dimention
		currPoint.Z++
		done = currPoint.Z >= limit.Z
	}


	return currPoint, done
}

func (node Node) MoveTo(space []space.Point, limit space.Point) ([]space.Point, bool) { // TODO: fix logic
	/**
	Node must first check if point exists in
	the occupied points Slice. If it does, it
	chooses a different point within the nearest space.
	*/

	// free previously occupied point
	
	if there, pIndex := node.IndexIn(space); there {
		// fmt.Printf("Node is there in space at index %d, now removing it\n", pIndex)
		space = slices.Delete(space, pIndex, pIndex+1)
		// there, _ := node.IndexIn(space)
		// fmt.Printf("There: %t\n", there)
	}

	var mu sync.Mutex

	var done bool
	mu.Lock()
	node.P, done = nextPoint(node.P, limit)
	mu.Unlock()

	there, _ := node.IndexIn(space)

	if there /* collision has occured */ {
		node.P = NextAvailableSpacePoint(space, node.P)
	}

	// fmt.Printf("pIndex = %d, pIndex + 1 = %d", pIndex, pIndex+1)
	// fmt.Printf("\tspace length = %d\n", len(space))

	return space, done
}

func NextAvailableSpacePoint(space []space.Point, p space.Point) space.Point {
	/*
		Finds the nearest point with the space of [0, 1] for any of the points as needed.
	*/
	availP := p

	if availP.X < availP.Y || availP.X < availP.Z {
		availP.X++
	} else if availP.Y < availP.X || availP.Y < availP.Z {
		availP.Y++
	} else {
		availP.Z++
	}

	

	for there := slices.Contains(space, availP) /* ?? May not work */ ; there == true ; {
		// run to find next available point as long no free point i.e., !there, is found
		NextAvailableSpacePoint(space, availP)
	}

	return availP // Not exactly useful, func just needs to return. Actual return statement at the top
}
