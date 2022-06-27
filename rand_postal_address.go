package fake

import "github.com/ejuju/fake/internal/geo"

func PostalAddress(addressGenerator geo.AddressGenerator) geo.Address {
	return addressGenerator.Generate()
}

func PostalAddressString(addressGenerator geo.AddressGenerator) string {
	return addressGenerator.String()
}
