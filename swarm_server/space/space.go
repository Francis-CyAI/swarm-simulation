package space

type Point struct { X, Y, Z int16 }

func (point Point) IndexIn(space []Point) (bool, int) {
	there := false
	i := -1
	for j := 0; j < len(space); j++ {
		if currSPoint := space[j]; point.X == currSPoint.X && point.Y == currSPoint.Y && point.Z == currSPoint.Z {
			there = true
			i = j
		}
	}
	return there, i
}

type Plane struct {	Origin, End Point }

