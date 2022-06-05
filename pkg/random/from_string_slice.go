package random

import "math/rand"

func FromStringSlice(strSlice []string) string {
	return strSlice[rand.Intn(len(strSlice))]
}
