package functions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuclideanDistance(t *testing.T) {
	assert.Equal(t, math.Sqrt(2), Distance(
		"euclidean",
		[]float64{1.0, 1.0},
		[]float64{0.0, 0.0},
	))

	assert.Equal(t, 1.0, Distance(
		"euclidean",
		[]float64{0.0, 1.0},
		[]float64{0.0, 0.0},
	))

	assert.Equal(t, 1.0, Distance(
		"euclidean",
		[]float64{math.NaN(), 1.0},
		[]float64{0.5, 0.0},
	))
}

func TestManhattanDistance(t *testing.T) {
	assert.Equal(t, 2.0, Distance(
		"manhattan",
		[]float64{1.0, 1.0},
		[]float64{0.0, 0.0},
	))

	assert.Equal(t, 1.0, Distance(
		"manhattan",
		[]float64{0.0, 1.0},
		[]float64{0.0, 0.0},
	))

	assert.Equal(t, 1.0, Distance(
		"manhattan",
		[]float64{math.NaN(), 1.0},
		[]float64{0.5, 0.0},
	))
}

func TestCoolingFunctions(t *testing.T) {
	assert.True(t, CoolingFactor("linear", 0.0) > 0.95)
	assert.True(t, CoolingFactor("soft", 0.0) > 0.95)
	assert.True(t, CoolingFactor("medium", 0.0) > 0.95)
	assert.True(t, CoolingFactor("hard", 0.0) > 0.95)

	assert.True(t, CoolingFactor("linear", 0.5) > 0.0)
	assert.True(t, CoolingFactor("soft", 0.5) > 0.0)
	assert.True(t, CoolingFactor("medium", 0.5) > 0.0)
	assert.True(t, CoolingFactor("hard", 0.5) > 0.0)

	assert.True(t, CoolingFactor("linear", 1.0) < 0.5)
	assert.True(t, CoolingFactor("soft", 1.0) < 0.5)
	assert.True(t, CoolingFactor("medium", 1.0) < 0.5)
	assert.True(t, CoolingFactor("hard", 1.0) < 0.5)
}

func TestNeighborhoodFunctions(t *testing.T) {
	assert.True(t, NeighborhoodInfluence("bubble", 0.5) > 0)
	assert.True(t, NeighborhoodInfluence("cone", 0.5) > 0)
	assert.True(t, NeighborhoodInfluence("gaussian", 0.5) > 0)
	assert.True(t, NeighborhoodInfluence("epanechicov", 0.5) > 0)
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, 1, min(2, 1))
}
