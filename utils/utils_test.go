package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/manhattanite/models"
	"github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	logger.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestLoadServerPort(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	port := LoadServerPort()

	g.Expect(port).ToNot(gomega.Equal(""), "HTTP Server is not empty")
}

func TestLoadDataFilePath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	path := LoadDataFilePath()

	g.Expect(path).ToNot(gomega.Equal(""), "Data file path is not empty")
}

func TestCalculateManhattanDistance(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	queried := models.Point{
		X: 10,
		Y: 10,
	}
	loaded := models.Point{
		X: 20,
		Y: 20,
	}

	manhattanDistance := CalculateManhattanDistance(queried, loaded)
	g.Expect(manhattanDistance).To(gomega.Equal(20), "Manhattan distance is equal to 20 (|10 - 20| + |10 - 20|)")
}

func TestParseIntParameterMissingParameter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{
		"IntParameter": []string{"5"},
	}

	param, err := parseIntParameter("Other", query)
	g.Expect(param).To(gomega.Equal(0), "Returned parameter is 0")
	g.Expect(err).To(gomega.MatchError(ErrInvalidQuery("missing parameter Other")), "Error should be a ErrInvalidQuery")
}
func TestParseIntParameterInvalidParameter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{
		"InvalidParameter": []string{"A"},
	}

	param, err := parseIntParameter("InvalidParameter", query)
	g.Expect(param).To(gomega.Equal(0), "Returned parameter is 0")
	g.Expect(err).To(gomega.MatchError(ErrInvalidQuery("could not convert parameter InvalidParameter to int")), "Error should be a ErrInvalidQuery")
}

func TestParseIntParameter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{
		"IntParameter": []string{"5"},
	}

	param, err := parseIntParameter("IntParameter", query)
	g.Expect(err).To(gomega.BeNil(), "Error is nil")
	g.Expect(param).To(gomega.Equal(5), "Returned converted parameter is equal to 5")
}

func TestParseQueryParametersEmpty(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{}

	queryParams, err := parseQueryParameters(query)

	g.Expect(queryParams).To(gomega.BeNil(), "queryParams is nil")
	g.Expect(err).To(gomega.MatchError(ErrInvalidQuery("empty query")), "Error should be a ErrInvalidQuery")
}

func TestParseQueryParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{
		"distance": []string{"5"},
		"x":        []string{"10"},
		"y":        []string{"20"},
	}

	queryParams, err := parseQueryParameters(query)

	g.Expect(err).To(gomega.BeNil(), "Error is nil")
	g.Expect(queryParams.MaxManhattanDistance).To(gomega.Equal(5), "Returned queryParams MaxManhattanDistance is equal to 5")
	g.Expect(queryParams.X).To(gomega.Equal(10), "Returned queryParams X is equal to 10")
	g.Expect(queryParams.Y).To(gomega.Equal(20), "Returned queryParams X is equal to 20")
}

func TestValidateRequest(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	query := map[string][]string{
		"distance": []string{"5"},
		"x":        []string{"10"},
		"y":        []string{"20"},
	}

	queryParams, err := ValidateRequest(query)
	g.Expect(err).To(gomega.BeNil(), "Error is nil")
	g.Expect(queryParams.MaxManhattanDistance).To(gomega.Equal(5), "Returned queryParams MaxManhattanDistance is equal to 5")
	g.Expect(queryParams.X).To(gomega.Equal(10), "Returned queryParams X is equal to 10")
	g.Expect(queryParams.Y).To(gomega.Equal(20), "Returned queryParams X is equal to 20")
}

func TestValidateHTTPMethodInvalid(t *testing.T) {

	var methodTableTests = []struct {
		in  string
		out error
	}{
		{"HEAD", ErrInvalidMethod("HEAD")},
		{"OPTIONS", ErrInvalidMethod("OPTIONS")},
		{"OPTIONS", ErrInvalidMethod("CONNECT")},
		{"TRACE", ErrInvalidMethod("TRACE")},
		{"POST", ErrInvalidMethod("POST")},
		{"PUT", ErrInvalidMethod("PUT")},
		{"PATCH", ErrInvalidMethod("PATCH")},
		{"DELETE", ErrInvalidMethod("DELETE")},
	}

	for _, tt := range methodTableTests {
		t.Run(tt.in, func(t *testing.T) {
			g := gomega.NewGomegaWithT(t)

			err := ValidateHTTPMethod(tt.in)
			g.Expect(err).To(gomega.MatchError(ErrInvalidMethod(tt.in)), "Error should be a ErrInvalidQuery")
		})
	}
}

func TestValidateHTTPMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	err := ValidateHTTPMethod("GET")

	g.Expect(err).To(gomega.BeNil(), "Error is nil")
}
