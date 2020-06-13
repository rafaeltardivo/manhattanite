package serializers

import (
	"encoding/json"

	"github.com/manhattanite/services"
)

// Cartesian serializer interface definition
type CartesianSerializer interface {
	Decode(raw []byte) ([]services.Point, error)
}

// JSON implementation of cartesian serializer interface
type JSONSerializer struct{}

// Decodes raw byte array to an array of points
func (j *JSONSerializer) Decode(raw []byte) ([]services.Point, error) {
	var points []services.Point

	if err := json.Unmarshal(raw, &points); err != nil {
		return nil, ErrSerialize(err.Error())
	}
	return points, nil
}
