package mathsskill

import "math"

func Average(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

// function for calculating variance.
func Variance(data []float64) float64 {
	mean := Average(data)

	var squareDifsum float64
	for _, val := range data {
		diff := val - mean
		squareDifsum += diff * diff
	}
	variance := squareDifsum / float64(len(data))

	return variance
}

func StandardDeviation(data []float64) float64 {
	variance := Variance(data)
	return math.Sqrt(variance)
}

func CalculateRange(data []float64) (float64, float64) {
	mean := Average(data)
	stdDev := StandardDeviation(data)
	lowerLimit := mean - stdDev*2
	upperLimit := mean + stdDev*2
	return lowerLimit, upperLimit
}
