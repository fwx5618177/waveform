package lib

import (
	"math"

	"azul3d.org/engine/audio"
)

// SampleReduceFunc is a function which reduces a set of float64 audio samples
// into a single float64 value.
type SampleReduceFunc func(samples audio.Float64) float64

// RMSF64Samples is a SampleReduceFunc which calculates the root mean square
// of a slice of float64 audio samples, enabling the measurement of magnitude
// over the entire set of samples.
//
// Derived from: http://en.wikipedia.org/wiki/Root_mean_square.
func RMSF64Samples(samples audio.Float64) float64 {
	// Square and sum all input samples
	var sumSquare float64
	for i := range samples {
		sumSquare += math.Pow(samples.At(i), 2)
	}

	// Multiply squared sum by length of samples slice, return square root
	return math.Sqrt(sumSquare / float64(samples.Len()))
}
