package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/manhattanite/models"
	"github.com/manhattanite/serializers"
)

// Query parameters data
type QueryParameters struct {
	X                    int
	Y                    int
	MaxManhattanDistance int
}

// Loads file according to environtment variable.
func LoadDataFile(envFile string, key string) ([]byte, error) {
	if err := godotenv.Load(envFile); err != nil {
		return nil, ErrNotFound("env file", envFile)
	}

	path := os.Getenv(key)
	if path == "" {
		return nil, ErrNotFound("key", key)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, ErrFile(err.Error())
	}

	return raw, nil
}

// Calculates and returns manhattan distance between queried and loaded point
// Formula: |queried.X - loaded.X| + |queried.Y - loaded.Y|
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

// Parses and return parameter as an integer
func parseIntParameter(name string, query url.Values) (int, error) {
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

// Return QueryParameters structured data
func parseQueryParameters(query url.Values) (*QueryParameters, error) {
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

// Validates request query parameters
func ValidateRequest(r *http.Request) (*QueryParameters, error) {
	rawParams := r.URL.Query()
	queryParams, err := parseQueryParameters(rawParams)

	if err != nil {
		return nil, err
	}

	return queryParams, nil
}

// Validates request http method (only GET is allowed)
func ValidateHTTPMethod(r *http.Request) error {
	if r.Method != http.MethodGet {
		return ErrInvalidMethod(r.Method)
	}

	return nil
}

// Returns HTTP response as JSON
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, string(serializers.NewJSONFormat().Encode(data)))
}
