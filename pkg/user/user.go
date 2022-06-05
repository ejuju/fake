package user

import (
	"net/mail"

	"github.com/ejuju/fake/pkg/email"
	"github.com/ejuju/fake/pkg/name"
	"github.com/ejuju/fake/pkg/random"
	"github.com/ejuju/fake/pkg/uuid"
)

type User struct {
	EmailAddress mail.Address
	UUID         string
	FirstName    string
	LastName     string
	Age          int
}

type UserGenerator struct {
	config UserGeneratorConfig
}

type UserGeneratorConfig struct {
	UUIDGenerator         uuid.Generator
	EmailAddressGenerator email.AddressGenerator
}

func NewUserGenerator(config UserGeneratorConfig) *UserGenerator {
	return &UserGenerator{config: config}
}

func (ug *UserGenerator) Generate() User {
	return User{
		UUID:         ug.config.UUIDGenerator.Generate(),
		FirstName:    random.FromStringSlice(name.SampleEnglishFirstNames),
		LastName:     random.FromStringSlice(name.SampleEnglishLastNames),
		Age:          random.IntMinMax(18, 55),
		EmailAddress: ug.config.EmailAddressGenerator.Generate(),
	}
}
