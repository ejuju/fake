package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ejuju/fake"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	// re-use the same seed if you want to get the same output
	// rand.Seed(1)

	// generate all bunch of fake random data for demonstratio purposes
	var (
		firstName  = fake.FirstName(nil)
		lastName   = fake.LastName(nil)
		fullName   = fake.Name(nil, nil)
		emailAddr  = fake.EmailAddress(nil).Address
		ipAddr     = fake.IPAddress().IP.String()
		postalAddr = fake.PostalAddressString(&fake.FrenchAddressGenerator{})
		paragraph  = fake.Paragraph(nil)
		user       = fake.User(nil)
		userAgent  = fake.HTTPUserAgent(nil)
		url        = fake.URL(nil, nil).String()
	)

	fmt.Printf("First name: %#v\n", firstName)
	fmt.Printf("Last name: %#v\n", lastName)
	fmt.Printf("Full name: %#v\n", fullName)
	fmt.Printf("Email address: %#v\n", emailAddr)
	fmt.Printf("IP address: %#v\n", ipAddr)
	fmt.Printf("Postal address: %#v\n", postalAddr)
	fmt.Printf("Paragraph: %#v\n", paragraph)
	fmt.Printf("User agent: %#v\n", userAgent)
	fmt.Printf("URL: %#v\n", url)
	fmt.Printf("User: %#v\n", user)
}
