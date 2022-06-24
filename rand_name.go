package fake

import (
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

func RandFirstName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = sample.EnglishFirstNames
	}
	return random.FromStringSlice(possibleNames)
}

func RandLastName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = sample.EnglishLastNames
	}
	return random.FromStringSlice(possibleNames)
}

func RandName(possibleFirstNames, possibleLastNames []string) string {
	return RandFirstName(possibleFirstNames) + " " + RandLastName(possibleLastNames)
}
