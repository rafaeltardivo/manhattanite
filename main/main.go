package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/manhattanite/models"
	"github.com/manhattanite/services"
	"github.com/manhattanite/utils"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	godotenv.Load("../.env")
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	service, err := services.NewCartesian()
	if err != nil {
		logger.Error("could not load cartesian service")
		os.Exit(1)
	}

	http.HandleFunc("/api/points", func(w http.ResponseWriter, r *http.Request) {
		if err := utils.ValidateHTTPMethod(r.Method); err != nil {
			logger.Error(err.Error())
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusForbidden)
			return
		}

		params, err := utils.ValidateRequest(r.URL.Query())
		if err != nil {
			logger.Error(err.Error())
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		queriedPoint := models.Point{X: params.X, Y: params.Y}
		points, err := service.FindPointsWithinDistance(queriedPoint, params.MaxManhattanDistance)
		if err != nil {
			logger.Error(err.Error())
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusNotFound)
			return
		}
		logger.Info(fmt.Sprintf("sorting and returning %d point(s) found within distance", len(points)))
		sort.Sort(models.PointDistanceSorter(points))
		utils.JSONResponse(w, points, http.StatusOK)
	})

	port := utils.LoadServerPort()
	if port == "" {
		logger.Error("could not load server port")
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("server is ready to accept connections on port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
