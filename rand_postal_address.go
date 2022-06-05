package fake

import "github.com/ejuju/fake/internal/postaladdress"

func RandPostalAddress(addressGenerator postaladdress.AddressGenerator) postaladdress.Address {
	return addressGenerator.Generate()
}

func RandPostalAddressString(addressGenerator postaladdress.AddressGenerator) string {
	return addressGenerator.String()
}
