package services

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/manhattanite/models"
	"github.com/manhattanite/utils"
	"github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	logger.SetOutput(ioutil.Discard)
	utils.Logger.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestFindPointsWithinDistanceNoPoints(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	service, _ := NewCartesian()

	queried := models.Point{
		X: -9999999,
		Y: -9999999,
	}

	points, err := service.FindPointsWithinDistance(queried, 5)

	g.Expect(points).To(gomega.BeNil(), "points are nil")
	g.Expect(err).To(gomega.MatchError(utils.ErrNotFound("point(s)", "no point(s) within specified distance")), "Error should be a ErrInvalidQuery")
}

func TestFindPointsWithinDistance(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	service, _ := NewCartesian()

	queried := models.Point{
		X: -30,
		Y: -38,
	}

	points, err := service.FindPointsWithinDistance(queried, 5)

	g.Expect(err).To(gomega.BeNil(), "Error is nil")
	g.Expect(points[0].X).To(gomega.Equal(-30), "X is -30")
	g.Expect(points[0].Y).To(gomega.Equal(-38), "Y is -38")
}
