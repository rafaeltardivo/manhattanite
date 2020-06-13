package serializers

import "github.com/manhattanite/services"

type CartesianSerializer interface {
	Decode(raw []byte) ([]services.Point, error)
}
