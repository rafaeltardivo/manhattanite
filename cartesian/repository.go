package cartesian

// Cartesian repository interface definition
type CartesianRepository interface {
	loadPoints() []Point
}

// Repository implementation of CartesianRepository
type repository struct{}

// Loads and returns  predefined points
func (r *repository) loadPoints() {

}
