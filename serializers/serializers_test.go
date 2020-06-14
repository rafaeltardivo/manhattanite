package serializers

import (
	"testing"

	"github.com/manhattanite/models"
	"github.com/onsi/gomega"
)

func TestNewSerializer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	JSONSerializer := NewJSONFormat()

	_, isOne := JSONSerializer.(CartesianSerializer)

	g.Expect(isOne).To(gomega.BeTrue(), "JSONSerializer is-one CartesianSerializer")
}

func TestSerializerEncodeDecodeSinglePoint(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	JSONSerializer := NewJSONFormat()

	point := []models.Point{
		{
			X: 10,
			Y: 10,
		},
		{
			X: 11,
			Y: 11,
		},
		{
			X: 12,
			Y: 12,
		},
	}

	encoded := JSONSerializer.Encode(point)
	decoded, err := JSONSerializer.Decode(encoded)

	g.Expect(err).To(gomega.BeNil(), "Error is nil")
	g.Expect(decoded[0].X).To(gomega.Equal(point[0].X), "Decoded[0].X is equal to point[0].X")
	g.Expect(decoded[0].Y).To(gomega.Equal(point[0].Y), "Decoded[0].Y is equal to point[0].Y")
	g.Expect(decoded[1].X).To(gomega.Equal(point[1].X), "Decoded[1].X is equal to point[1].X")
	g.Expect(decoded[1].Y).To(gomega.Equal(point[1].Y), "Decoded[1].Y is equal to point[1].Y")
}
