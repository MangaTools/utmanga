package algorithm

import (
	"math"
	"sort"
)

var dotCentralPointOffset = point2D[float64]{
	X: .1,
	Y: .15,
}

var (
	defaultWhitePoint byte = 255
	defaultGrayPoint  byte = 128
	defaultBlackPoint byte = 1
)

// generate
func generateDot(size int, inverted bool) []byte {
	pointValues := generateDotPointValues(size)

	sortPoints(pointValues, inverted)

	if inverted {
		return generateThresholdByPointValues(pointValues, size, defaultBlackPoint, defaultGrayPoint)
	}

	return generateThresholdByPointValues(pointValues, size, defaultGrayPoint+1, defaultWhitePoint)
}

// generateDotPointValues creates PointValue2D slice with value equals distance from dot pixel to central point.
func generateDotPointValues(size int) []pointValue2D[int, float64] {
	centralValue := float64(size-1) / 2
	centerPoint := point2D[float64]{X: centralValue + dotCentralPointOffset.X, Y: centralValue + dotCentralPointOffset.Y}

	valuePoints := make([]pointValue2D[int, float64], 0, size*size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// more value if point is farther from central point
			value := centerPoint.Distance(float64(x), float64(y))
			point := pointValue2D[int, float64]{
				point2D: point2D[int]{
					X: x,
					Y: y,
				},
				Value: value,
			}

			valuePoints = append(valuePoints, point)
		}
	}

	return valuePoints
}

// sortPoints according to value.
func sortPoints(points []pointValue2D[int, float64], inverted bool) {
	if inverted {
		sort.Slice(points, func(i, j int) bool { return !(points[i].Value < points[j].Value) })
	} else {
		sort.Slice(points, func(i, j int) bool { return points[i].Value < points[j].Value })
	}
}

// generateThresholdByPointValues return matrix with threshold. values must be sorted
func generateThresholdByPointValues(values []pointValue2D[int, float64], size int, min, max byte) []byte {
	matrix := make([]byte, len(values))
	// NOTE(shadream): example: min = 1, max = 3. points are 2. first point = 1, second point = 3. step is equal 2. (3-1)/(2-x) = 2, x is 1.
	stepValue := float64(max-min) / float64(len(values)-1)

	for i, value := range values {
		newValue := math.Round(float64(max) - (stepValue * float64(i)))
		matrix[value.X+(value.Y*size)] = byte(newValue)
	}

	return matrix
}
