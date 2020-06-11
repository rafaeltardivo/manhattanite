package cartesian

type Point struct {
	x int
	y int
}

// Cartesian service interface definition
type CartesianService interface {
	findPointsWithinDistance(Point, int)
}

// Service implementation according to CartesianService
type service struct{}

// Finds and returns points within queried point origin and distance
func (s *service) findPointsWithinDistance(queriedPoint Point, maxDistance int) ([]Point, error) {
	return nil, nil
}
