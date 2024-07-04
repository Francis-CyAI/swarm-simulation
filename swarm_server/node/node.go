package node

import (
	"fmt" // TODO: remove
	"net"
	"slices"
	"swarm_server/space"
	"sync"
)

type Node struct {
	R float32     /// Radius
	P space.Point /// Position of center
}

func (node Node) IndexIn(space []space.Point) (bool, int) { // Works fine
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

func nextPoint(currPoint space.Point, limit space.Point) (space.Point, bool) {
	//  Okay: Increments fine, and gives the right done status.

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

func (node *Node) MoveTo(c net.Conn, space []space.Point, limit space.Point) (space.Point, bool) { // TODO: fix logic
	// Okay: working perfectly

	/**
	Node must first check if point exists in
	the occupied points Slice. If it does, it
	chooses a different point within the nearest space.
	*/

	var mu sync.Mutex

	var done bool

	for {

		// free previously occupied point
		if there, pIndex := node.IndexIn(space); there {
			// fmt.Printf("Node is there in space at index %d, now removing it\n", pIndex)
			space = slices.Delete(space, pIndex, pIndex+1)
			// there, _ := node.IndexIn(space)
			// fmt.Printf("There: %t\n", there)
		}

		mu.Lock()
		node.P, done = nextPoint(node.P, limit)
		mu.Unlock()

		// Report progress
		fmt.Fprintf(c, "nextPoint:- X: %d, Y: %d, Z: %d; done: %t\n", node.P.X, node.P.Y, node.P.Z, done)

		there, _ := node.IndexIn(space)

		if there /* collision has occured */ {
			fmt.Fprintln(c, "Point already occupied")
			notFound := true
			for {
				node.P, notFound = NextAvailableSpacePoint(space, node.P, limit)
				if !notFound {
					fmt.Fprintf(c, "Free space point found:- X: %d, Y: %d, Z: %d; done: %t\n", node.P.X, node.P.Y, node.P.Z, done)

					break
				} // A free space point has been found
				//TODO: if space vicinity exhausted, move to 'nextPoint'
			}
		}

		// fmt.Printf("pIndex = %d, pIndex + 1 = %d", pIndex, pIndex+1)
		// fmt.Printf("\tspace length = %d\n", len(space))
		if done {
			break
		}
	}
	c.Close()
	return node.P, done
}

func NextAvailableSpacePoint(space []space.Point, p space.Point, limit space.Point) (space.Point, bool) {
	/*
		Finds the nearest point with the limit for any of the secondary axes
	*/

	notFound := true
	availP := p

	if availP.X > availP.Y && availP.X > availP.Z {
		// Exclude X as it is the main axis
		if availP.Y < availP.Z && availP.Y < limit.Y {
			availP.Y++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else if availP.Z < limit.Z {
			availP.Z++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else {
			// not available space point found
			return p, notFound
		}
	} else if availP.Y > availP.X && availP.Y > availP.Z {
		// Exclude Y as it is the main axis
		if availP.X < availP.Z && availP.X < limit.X {
			availP.X++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else if availP.Z < limit.Z {
			availP.Z++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else {
			return p, notFound
		}
	} else if availP.Z > availP.X && availP.Z > availP.Y {
		// Exclude Z as it is the main axis
		if availP.X < availP.Y && availP.X < limit.X {
			availP.X++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else if availP.Y < limit.Y {
			availP.Y++
			if occupied, _ := indexIn(space, availP); occupied {
				NextAvailableSpacePoint(space, availP, limit)
			}
		} else {
			return p, notFound
		}
	}

	/*
		for there := slices.Contains(space, availP) /* ?? May not work / ; there == true; {
			fmt.Printf("availP X = %d, Y = %d, Z = %d; there = %t", availP.X, availP.Y, availP.Z, there)
			// run to find next available point as long no free point i.e., !there, is found
			NextAvailableSpacePoint(space, availP)
		}
	*/

	return availP, !notFound // Found a free space point
}
