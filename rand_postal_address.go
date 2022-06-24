package fake

import "github.com/ejuju/fake/internal/postaladdress"

func PostalAddress(addressGenerator postaladdress.AddressGenerator) postaladdress.Address {
	return addressGenerator.Generate()
}

func PostalAddressString(addressGenerator postaladdress.AddressGenerator) string {
	return addressGenerator.String()
}
