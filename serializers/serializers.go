package serializers

import (
	"encoding/json"

	"github.com/manhattanite/models"
)

// Cartesian serializer interface definition
type CartesianSerializer interface {
	Decode(raw []byte) ([]models.Point, error)
}

// JSON implementation of cartesian serializer interface
type JSONFormat struct{}

// Decodes raw byte array to an array of points
func (j *JSONFormat) Decode(raw []byte) ([]models.Point, error) {
	var points []models.Point

	if err := json.Unmarshal(raw, &points); err != nil {
		return nil, ErrSerialize(err.Error())
	}
	return points, nil
}

// Returns a new JSONFormat.
func NewJSONFormat() CartesianSerializer {
	return &JSONFormat{}
}
