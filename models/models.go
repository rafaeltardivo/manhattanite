package models

// Cartesian point model.
// Ignores ManhattanDistance for JSON serialization.
type Point struct {
	X                 int `json:"x"`
	Y                 int `json:"y"`
	ManhattanDistance int `json:"-"`
}

// Point implementation of native sort interface.
type PointDistanceSorter []Point

// Returns size of point slice (implementation of len).
func (p PointDistanceSorter) Len() int {
	return len(p)
}

// Returns comparison between current and next element (implementation of Less).
func (p PointDistanceSorter) Less(i, j int) bool {
	return p[i].ManhattanDistance < p[j].ManhattanDistance
}

// Swaps current for next element(implementation of Swap).
func (p PointDistanceSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Returns a new Point model.
func NewPoint(p Point, manhattanDistance int) Point {
	return Point{
		X:                 p.X,
		Y:                 p.Y,
		ManhattanDistance: manhattanDistance,
	}
}
