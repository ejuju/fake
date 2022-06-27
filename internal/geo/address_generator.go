package geo

// AddressGenerator can generate an address in the right format of a given country
type AddressGenerator interface {
	Generate() Address
	String() string
}

// Address holds information about a postal address
type Address struct {
	CountryName  string
	CountryCode  string // abbreviation
	ZipCode      string // postal code
	CityName     string
	StreetName   string
	StreetNumber string
}
