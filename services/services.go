package services

import (
	"fmt"

	"github.com/manhattanite/models"
	"github.com/manhattanite/repositories"
	"github.com/manhattanite/utils"
	"github.com/sirupsen/logrus"
)

// Cartesian service interface definition.
type CartesianService interface {
	FindPointsWithinDistance([]models.Point, int) ([]models.Point, error)
}

// Service implementation according to CartesianService.
type cartesian struct {
	points []models.Point
}

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.JSONFormatter{})
}

// Finds and returns points within queried point manhattahn distance.
func (c *cartesian) FindPointsWithinDistance(queried models.Point, maxDistance int) ([]models.Point, error) {
	var pointsWithinDistance []models.Point

	logger.Info(fmt.Sprintf("querying for points within distance x=%d y=%d distance=%d ", queried.X, queried.Y, maxDistance))
	for _, point := range c.points {
		manhattanDistance := utils.CalculateManhattanDistance(queried, point)
		if manhattanDistance <= maxDistance {
			pointsWithinDistance = append(pointsWithinDistance, models.NewPoint(point, manhattanDistance))
		}
	}

	if len(pointsWithinDistance) < 1 {
		return nil, utils.ErrNotFound("point(s)", "no point(s) within specified distance")
	}

	return pointsWithinDistance, nil
}

// Returns a new cartesian service (with loaded points).
func NewCartesian() (*cartesian, error) {
	local := repositories.NewLocalJSONFile()
	loadedPoints, err := local.LoadPoints()

	if err != nil {
		return nil, err
	}

	return &cartesian{
		points: loadedPoints,
	}, nil
}
