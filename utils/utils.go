package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/manhattanite/models"
	"github.com/manhattanite/serializers"
	"github.com/sirupsen/logrus"
)

// Query parameters data
type QueryParameters struct {
	X                    int
	Y                    int
	MaxManhattanDistance int
}

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.JSONFormatter{})
}

// Loads server port from to environment variable.
func LoadServerPort() string {
	key := "HTTP_SERVER_PORT"

	Logger.Info(fmt.Sprintf("loading key %s from environment", key))
	return os.Getenv(key)
}

// Loads data file path from environment variable.
func LoadDataFilePath() string {
	key := "POINTS_FILE_RELATIVE_PATH"

	Logger.Info(fmt.Sprintf("loading key %s from environment", key))
	return os.Getenv(key)
}

// Loads file according to environtment variable.
func LoadDataFile() ([]byte, error) {
	path := LoadDataFilePath()
	if path == "" {
		return nil, ErrNotFound("key", "could not load file path")
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, ErrFile(err.Error())
	}

	return raw, nil
}

// Calculates and returns manhattan distance between queried and loaded point.
// Formula: |queried.X - loaded.X| + |queried.Y - loaded.Y|.
func CalculateManhattanDistance(queried models.Point, loaded models.Point) int {
	x := queried.X - loaded.X
	// Absolute value
	if x < 0 {
		x *= -1
	}
	y := queried.Y - loaded.Y
	// Absolute value
	if y < 0 {
		y *= -1
	}

	return x + y
}

// Parses and return parameter as an integer.
func parseIntParameter(name string, query map[string][]string) (int, error) {
	param, found := query[name]
	if !found {
		return 0, ErrInvalidQuery(fmt.Sprintf("missing parameter %s", name))
	}

	value, err := strconv.Atoi(param[0])
	if err != nil {
		return 0, ErrInvalidQuery(fmt.Sprintf("could not convert parameter %s to int", name))
	}

	return value, nil
}

// Parses and returns query parameters.
// Expected: x, y and distance.
func parseQueryParameters(query map[string][]string) (*QueryParameters, error) {
	if len(query) < 1 {
		return nil, ErrInvalidQuery("empty query")
	}

	x, err := parseIntParameter("x", query)
	if err != nil {
		return nil, err
	}
	y, err := parseIntParameter("y", query)
	if err != nil {
		return nil, err
	}
	distance, err := parseIntParameter("distance", query)
	if err != nil {
		return nil, err
	}

	return &QueryParameters{
		X:                    x,
		Y:                    y,
		MaxManhattanDistance: distance,
	}, nil
}

// Validates and returns request query parameters.
func ValidateRequest(rawParams map[string][]string) (*QueryParameters, error) {
	Logger.Info(fmt.Sprintf("validating request parameters %s", rawParams))
	queryParams, err := parseQueryParameters(rawParams)

	if err != nil {
		return nil, err
	}

	return queryParams, nil
}

// Validates request http method (only GET is allowed).
func ValidateHTTPMethod(HTTPMethod string) error {
	if HTTPMethod != http.MethodGet {
		return ErrInvalidMethod(HTTPMethod)
	}

	return nil
}

// Returns HTTP response as JSON.
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	fmt.Fprintln(w, string(serializers.NewJSONFormat().Encode(data)))
}
