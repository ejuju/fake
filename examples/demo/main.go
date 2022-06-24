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
		firstName  = fake.RandFirstName(nil)
		lastName   = fake.RandLastName(nil)
		fullName   = fake.RandName(nil, nil)
		emailAddr  = fake.RandEmailAddress(nil).Address
		ipAddr     = fake.RandIPAddress().IP.String()
		postalAddr = fake.RandPostalAddressString(&fake.FrenchAddressGenerator{})
		paragraph  = fake.RandParagraph(nil)
		user       = fake.RandUser(nil)
		userAgent  = fake.RandUserAgent(nil)
		url        = fake.RandURL(nil, nil)
	)

	fmt.Printf("First name: %s\n", firstName)
	fmt.Printf("Last name: %s\n", lastName)
	fmt.Printf("Full name: %s\n", fullName)
	fmt.Printf("Email address: %s\n", emailAddr)
	fmt.Printf("IP address: %s\n", ipAddr)
	fmt.Printf("Postal address: %s\n", postalAddr)
	fmt.Printf("Paragraph: %+v\n", paragraph)
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("User agent: %+v\n", userAgent)
	fmt.Printf("URL: %+v\n", url)
}
