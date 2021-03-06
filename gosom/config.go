package main

import (
	"strconv"

	"github.com/docopt/docopt-go"
)

type config struct {
	prepare     bool
	train       bool
	classify    bool
	interpolate bool
	plot        bool
	test        bool
	functions   bool

	file                 string
	directory            string
	data                 string
	width                int
	height               int
	initialization       string
	distanceFunction     string
	neighborhoodFunction string
	coolingFunction      string
	trainingSteps        int
	initialLearningRate  float64
	finalLearningRate    float64
	initialRadius        float64
	finalRadius          float64
	input                string
	weighted             bool
	nearestNeighbors     int
	size                 int
	prefix               string
	testDimensions       int
	quiet                bool
}

func parseConfig() *config {
	usage := `Self organizing maps for go.

Usage:
  gosom prepare <file> <data> <width> <height> [-i <im> -d <df> -n <nf> -c <cf>]
  gosom train <file> <data> [-t <ts> -l <lr> -m <lr> -r <nr> -g <nr>]
  gosom classify <file> <input>
  gosom interpolate <file> <input> [-w -k <nn>]
  gosom plot <file> <directory> [-s <ns> -p <fp>]
  gosom test <file> <data> [-k <nn> -j <td> -q]
  gosom -f
  gosom -h
  gosom -v

Options:
  -i <im>  Initialization method (random, datapoints) [default: datapoints].
  -d <df>  Distance function (euclidean, manhattan) [default: euclidean].
  -n <nf>  Neighborhood function (bubble, cone, gaussian, epanechicov) [default: cone].
  -c <cf>  Cooling function (linear, soft, medium, hard) [default: linear].
  -t <ts>  Number of training steps [default: 10000].
  -l <lr>  Initial learning rate [default: 0.5].
  -m <lr>  Final learning rate [default: 0.05].
  -r <nr>  Initial neighborhood radius [default: -1].
  -g <nr>  Final neighborhood radius [default: 0.0].
  -k <nn>  Number of nearest neighbors to consider [default: 5].
  -w       Use weighted interpolation.
  -s <ns>  Size of the individual nodes [default: 10].
  -p <fp>  Filename prefix [default: som].
  -j <td>  Number of dimensions to test [default: 1].
  -q       Quiet output.
  -f       Plot functions to current directoy.
  -h       Show help.
  -v       Show version.`

	a, err := docopt.Parse(usage, nil, true, "gosom 0.1", false)
	if err != nil {
		panic(err)
	}

	return &config{
		prepare:              getBool(a["prepare"]),
		train:                getBool(a["train"]),
		classify:             getBool(a["classify"]),
		interpolate:          getBool(a["interpolate"]),
		plot:                 getBool(a["plot"]),
		test:                 getBool(a["test"]),
		functions:            getBool(a["-f"]),
		file:                 getString(a["<file>"]),
		directory:            getString(a["<directory>"]),
		data:                 getString(a["<data>"]),
		width:                getInt(a["<width>"]),
		height:               getInt(a["<height>"]),
		initialization:       getString(a["-i"]),
		distanceFunction:     getString(a["-d"]),
		neighborhoodFunction: getString(a["-n"]),
		coolingFunction:      getString(a["-c"]),
		trainingSteps:        getInt(a["-t"]),
		initialLearningRate:  getFloat(a["-l"]),
		finalLearningRate:    getFloat(a["-m"]),
		initialRadius:        getFloat(a["-r"]),
		finalRadius:          getFloat(a["-g"]),
		input:                getString(a["<input>"]),
		weighted:             getBool(a["-w"]),
		nearestNeighbors:     getInt(a["-k"]),
		size:                 getInt(a["-s"]),
		prefix:               getString(a["-p"]),
		testDimensions:       getInt(a["-j"]),
		quiet:                getBool(a["-q"]),
	}
}

func getBool(v interface{}) bool {
	b, ok := v.(bool)

	if !ok {
		return false
	}

	return b
}

func getString(v interface{}) string {
	s, ok := v.(string)

	if !ok {
		return ""
	}

	return s
}

func getInt(v interface{}) int {
	s := getString(v)

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

func getFloat(v interface{}) float64 {
	s := getString(v)

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}

	return f
}
