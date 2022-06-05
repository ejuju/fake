package random

import "math/rand"

// The minimum is inclusive and maximum is exclusive. [min;max[
func IntMinMax(min, max int) int {
	return min + rand.Intn(max-min)
}
