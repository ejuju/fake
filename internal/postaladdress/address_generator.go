package postaladdress

type AddressGenerator interface {
	Generate() Address
	String() string
}

type Address struct {
	CountryName  string
	CountryCode  string
	ZipCode      string
	CityName     string
	StreetName   string
	StreetNumber string
}
