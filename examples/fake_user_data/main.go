package main

import (
	"fmt"

	"github.com/ejuju/fake/pkg/email"
	"github.com/ejuju/fake/pkg/name"
	"github.com/ejuju/fake/pkg/url"
	"github.com/ejuju/fake/pkg/user"
	"github.com/ejuju/fake/pkg/uuid"
)

func main() {
	usergen := user.NewUserGenerator(user.UserGeneratorConfig{
		UUIDGenerator: uuid.NewGoogleUUIDGenerator(),
		EmailAddressGenerator: *email.NewEmailAddressGenerator(email.AddressGeneratorConfig{
			FirstNames:               name.SampleEnglishFirstNames,
			LastNames:                name.SampleEnglishLastNames,
			EmailProviderDomainNames: url.SampleEmailProviderDomainNames,
		}),
	})

	for i := 0; i < 5; i++ {
		fmt.Printf("User %d: %+v\n", i, usergen.Generate())
	}
}
