package fake

import "github.com/ejuju/fake/internal/random"

func RandFirstName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = SampleEnglishFirstNames
	}
	return random.FromStringSlice(possibleNames)
}

func RandLastName(possibleNames []string) string {
	if len(possibleNames) == 0 {
		possibleNames = SampleEnglishLastNames
	}
	return random.FromStringSlice(possibleNames)
}

func RandName(possibleFirstNames, possibleLastNames []string) string {
	return RandFirstName(possibleFirstNames) + " " + RandLastName(possibleLastNames)
}
