package main

import (
	"fmt"

	"github.com/256dpi/gosom"
)

var data = [][]float64{
	{ 0.0, 4.0 },
	{ 1.0, 3.0 },
	{ 2.0, 2.0 },
	{ 3.0, 1.0 },
	{ 4.0, 0.0 },
}

func main() {
	som := gosom.NewSOM(data, 8, 8)
	som.InitializeWithRandomDataPoints()
	som.DistanceFunction = gosom.ManhattanDistance
	som.CoolingFunction = gosom.MediumCooling
	som.NeighborhoodFunction = gosom.GaussianNeighborhood

	som.Train(10000, 0.5)
	fmt.Println(som)

	fmt.Printf("3.5: %f\n", som.Classify([]float64{0.5}))
	fmt.Printf("2.5: %f\n", som.Classify([]float64{1.5}))
	fmt.Printf("1.5: %f\n", som.Classify([]float64{2.5}))
	fmt.Printf("0.5: %f\n\n", som.Classify([]float64{3.5}))

	fmt.Printf("3.5: %f\n", som.Interpolate([]float64{0.5}, 16))
	fmt.Printf("2.5: %f\n", som.Interpolate([]float64{1.5}, 16))
	fmt.Printf("1.5: %f\n", som.Interpolate([]float64{2.5}, 16))
	fmt.Printf("0.5: %f\n", som.Interpolate([]float64{3.5}, 16))
}
