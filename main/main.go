package main

import (
	"net/http"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/manhattanite/models"
	"github.com/manhattanite/services"
	"github.com/manhattanite/utils"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	service, err := services.NewCartesian()
	if err != nil {
		// TODO log
		os.Exit(1)
	}

	http.HandleFunc("/api/points", func(w http.ResponseWriter, r *http.Request) {

		if err := utils.ValidateHTTPMethod(r.Method); err != nil {
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusForbidden)
			return
		}

		params, err := utils.ValidateRequest(r.URL.Query())
		if err != nil {
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		queriedPoint := models.Point{X: params.X, Y: params.Y}
		points, err := service.FindPointsWithinDistance(queriedPoint, params.MaxManhattanDistance)
		if err != nil {
			utils.JSONResponse(w, utils.ErrorResponse{Msg: err.Error()}, http.StatusNotFound)
			return
		}
		sort.Sort(models.PointDistanceSorter(points))
		utils.JSONResponse(w, points, http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
