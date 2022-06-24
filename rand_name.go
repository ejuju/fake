package fake

import (
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

func FirstName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = sample.EnglishFirstNames
	}
	return random.FromStringSlice(possibleNames)
}

func LastName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = sample.EnglishLastNames
	}
	return random.FromStringSlice(possibleNames)
}

func Name(possibleFirstNames, possibleLastNames []string) string {
	return FirstName(possibleFirstNames) + " " + LastName(possibleLastNames)
}
