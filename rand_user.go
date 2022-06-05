package fake

import (
	"net/mail"

	"github.com/ejuju/fake/internal/uuid"
)

type User struct {
	EmailAddress mail.Address `json:"email_address"`
	ID           string       `json:"id"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
}

type RandUserConfig struct {
	PossibleFirstNames     []string
	PossibleLastNames      []string
	PossibleEmailProviders []string
	MinAge                 int
	MaxAge                 int
}

func RandUser(config *RandUserConfig) User {
	if config == nil {
		config = &RandUserConfig{MinAge: 18}
	}
	if config.MaxAge == 0 {
		config.MaxAge = 60
	}

	return User{
		ID:        uuid.NewUUIDWithFallback(),
		FirstName: RandFirstName(config.PossibleFirstNames),
		LastName:  RandLastName(config.PossibleLastNames),
		EmailAddress: RandEmailAddress(&RandEmailAddressConfig{
			FirstNameList:               config.PossibleFirstNames,
			LastNameList:                config.PossibleLastNames,
			EmailProviderDomainNameList: config.PossibleEmailProviders,
		}),
	}
}
