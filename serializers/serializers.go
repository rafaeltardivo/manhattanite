package serializers

import (
	"encoding/json"

	"github.com/manhattanite/models"
)

// CartesianSerializer interface definition.
type CartesianSerializer interface {
	Decode([]byte) ([]models.Point, error)
	Encode(interface{}) []byte
}

// JSONFormat implementation of cartesian serializer interface.
type JSONFormat struct{}

// Decode decodes raw byte array to an array of points.
func (j *JSONFormat) Decode(raw []byte) ([]models.Point, error) {
	var points []models.Point

	if err := json.Unmarshal(raw, &points); err != nil {
		return nil, ErrSerialize(err.Error())
	}
	return points, nil
}

// Encode encodes an interface into an array of bytes.
func (j *JSONFormat) Encode(data interface{}) []byte {
	payload, _ := json.Marshal(data)
	return payload
}

// NewJSONFormat() returns a new JSONFormat.
func NewJSONFormat() CartesianSerializer {
	return &JSONFormat{}
}
