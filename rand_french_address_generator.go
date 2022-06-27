package fake

import (
	"github.com/ejuju/fake/internal/geo"
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

type FrenchAddressGenerator struct {
	FrenchCities  []string
	FrenchStreets []string
	PostalCodes   []string
	StreetNumbers []string
	CountryName   string
	CountryCode   string
}

func (frenchAddressGenerator *FrenchAddressGenerator) Generate() geo.Address {
	if len(frenchAddressGenerator.FrenchCities) == 0 {
		frenchAddressGenerator.FrenchCities = sample.FrenchCities
	}
	if len(frenchAddressGenerator.FrenchStreets) == 0 {
		frenchAddressGenerator.FrenchStreets = sample.FrenchStreetNames
	}
	if len(frenchAddressGenerator.PostalCodes) == 0 {
		frenchAddressGenerator.PostalCodes = sample.FrenchPostalCodes
	}
	if len(frenchAddressGenerator.StreetNumbers) == 0 {
		frenchAddressGenerator.StreetNumbers = sample.FrenchStreetNumbers
	}
	if frenchAddressGenerator.CountryName == "" {
		frenchAddressGenerator.CountryName = "France"
	}
	if frenchAddressGenerator.CountryCode == "" {
		frenchAddressGenerator.CountryCode = "FR"
	}

	return geo.Address{
		CountryName:  frenchAddressGenerator.CountryName,
		CountryCode:  frenchAddressGenerator.CountryCode,
		ZipCode:      random.FromStringSlice(frenchAddressGenerator.PostalCodes),
		CityName:     random.FromStringSlice(frenchAddressGenerator.FrenchCities),
		StreetName:   random.FromStringSlice(frenchAddressGenerator.FrenchStreets),
		StreetNumber: random.FromStringSlice(frenchAddressGenerator.StreetNumbers),
	}
}

func (frenchAddressGenerator *FrenchAddressGenerator) String() string {
	addr := frenchAddressGenerator.Generate()

	// French street address format example:
	// "15 rue de la Garenne, 92500 Rueil-Malmaison, France"
	return addr.StreetNumber + " " +
		addr.StreetName + ", " +
		addr.ZipCode + " " +
		addr.CityName + ", " +
		addr.CountryName
}
