package fake

import (
	"strconv"

	"github.com/ejuju/fake/internal/postaladdress"
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

type FrenchAddressGenerator struct {
	FrenchCities  []string
	FrenchStreets []string
}

func (frenchAddressGenerator *FrenchAddressGenerator) Generate() postaladdress.Address {
	if len(frenchAddressGenerator.FrenchCities) == 0 {
		frenchAddressGenerator.FrenchCities = sample.FrenchCities
	}
	if len(frenchAddressGenerator.FrenchStreets) == 0 {
		frenchAddressGenerator.FrenchStreets = sample.FrenchStreetNames
	}

	return postaladdress.Address{
		CountryName:  "France",
		CountryCode:  "FR",
		ZipCode:      strconv.Itoa(random.IntMinMax(10000, 95700)),
		CityName:     random.FromStringSlice(frenchAddressGenerator.FrenchCities),
		StreetName:   random.FromStringSlice(frenchAddressGenerator.FrenchStreets),
		StreetNumber: strconv.Itoa(random.IntMinMax(1, 150)),
	}
}

func (frenchAddressGenerator *FrenchAddressGenerator) String() string {
	addr := frenchAddressGenerator.Generate()

	// French street address format example:
	// "15 rue de la Garenne, 92500 Rueil-Malmaison, France"
	return addr.StreetNumber + " " + addr.StreetName + ", " + addr.ZipCode + " " + addr.CityName + ", " + addr.CountryName
}
