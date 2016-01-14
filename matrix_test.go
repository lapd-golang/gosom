package gosom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var slice = [][]float64{
	{1.0, 0.5, 0.0},
	{0.0, 0.5, 1.0},
}

func TestMatrix(t *testing.T) {
	m := NewMatrix(slice)

	assert.Equal(t, slice, m.Data)
	assert.Equal(t, 2, m.Rows)
	assert.Equal(t, 3, m.Columns)
	assert.Equal(t, []float64{0.0, 0.5, 0.0}, m.Minimums)
	assert.Equal(t, []float64{1.0, 0.5, 1.0}, m.Maximums)
	assert.Equal(t, 0.0, m.Minimum)
	assert.Equal(t, 1.0, m.Maximum)
}

func TestSubMatrix1(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(0, 2)

	d := [][]float64{
		{1.0, 0.5},
		{0.0, 0.5},
	}

	assert.Equal(t, d, sm.Data)
}

func TestSubMatrix2(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(2, 1)

	d := [][]float64{
		{0.0},
		{1.0},
	}

	assert.Equal(t, d, sm.Data)
}

func TestRandomRow(t *testing.T) {
	m := NewMatrix(slice)

	t1 := assert.ObjectsAreEqual(m.RandomRow(), slice[0])
	t2 := assert.ObjectsAreEqual(m.RandomRow(), slice[1])

	assert.True(t, t1 || t2)
}

func TestLoadMatrixFromCSV(t *testing.T) {
	csv := "1.0,0.5,0.0\n0.0,0.5,1.0"
	reader := strings.NewReader(csv)

	m, err := LoadMatrixFromCSV(reader)

	assert.NoError(t, err)
	assert.Equal(t, slice, m.Data)
}

func TestLoadMatrixFromJSON(t *testing.T) {
	json := "[[1.0,0.5,0.0],[0.0,0.5,1.0]]"
	reader := strings.NewReader(json)

	m, err := LoadMatrixFromJSON(reader)

	assert.NoError(t, err)
	assert.Equal(t, slice, m.Data)

}
