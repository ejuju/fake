package fake

import (
	"strconv"

	"github.com/ejuju/fake/internal/postaladdress"
	"github.com/ejuju/fake/internal/random"
)

type FrenchAddressGenerator struct {
	FrenchCities  []string
	FrenchStreets []string
}

func (frenchAddressGenerator *FrenchAddressGenerator) Generate() postaladdress.Address {
	if len(frenchAddressGenerator.FrenchCities) == 0 {
		frenchAddressGenerator.FrenchCities = SampleFrenchCities
	}
	if len(frenchAddressGenerator.FrenchStreets) == 0 {
		frenchAddressGenerator.FrenchStreets = SampleFrenchStreetNames
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
	return addr.StreetNumber + " " + addr.StreetName + ", " + addr.ZipCode + " " + addr.CityName + ", " + addr.CountryName
}
