package fake

import (
	"net/mail"

	"github.com/ejuju/fake/internal/uuid"
)

type DefaultUser struct {
	ID           string       `json:"id"`
	EmailAddress mail.Address `json:"email_address"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
}

type UserConfig struct {
	PossibleFirstNames     []string
	PossibleLastNames      []string
	PossibleEmailProviders []string
	MinAge                 int
	MaxAge                 int
}

func User(config *UserConfig) DefaultUser {
	if config == nil {
		config = &UserConfig{MinAge: 18}
	}
	if config.MaxAge == 0 {
		config.MaxAge = 60
	}

	firstName := FirstName(config.PossibleFirstNames)
	lastName := LastName(config.PossibleLastNames)

	return DefaultUser{
		ID:        uuid.NewUUIDWithFallback(),
		FirstName: firstName,
		LastName:  lastName,
		EmailAddress: EmailAddress(&EmailAddressConfig{
			FirstNameList:               []string{firstName},
			LastNameList:                []string{lastName},
			EmailProviderDomainNameList: config.PossibleEmailProviders,
		}),
	}
}
