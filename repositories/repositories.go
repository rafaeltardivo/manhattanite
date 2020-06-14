package repositories

import (
	"github.com/manhattanite/models"
	"github.com/manhattanite/serializers"
	"github.com/manhattanite/utils"
)

// Cartesian repository interface definition
type CartesianRepository interface {
	LoadPoints() ([]models.Point, error)
}

// Local JSON file implementation of CartesianRepository.
type localJSONFile struct{}

// Loads and returns points data from json file.
func (l *localJSONFile) LoadPoints() ([]models.Point, error) {
	raw, err := utils.LoadDataFile("../.env", "POINTS_FILE_RELATIVE_PATH")
	if err != nil {
		return nil, err
	}
	return serializers.NewJSONFormat().Decode(raw)
}

// Returns a new localJSONFile repository.
func NewLocalJSONFile() CartesianRepository {
	return &localJSONFile{}
}
