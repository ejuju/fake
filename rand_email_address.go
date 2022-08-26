package fake

import (
	"net/mail"
	"strings"

	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/pkg/sample"
)

type EmailAddressConfig struct {
	FirstNameList               []string
	LastNameList                []string
	EmailProviderDomainNameList []string
}

func EmailAddress(config *EmailAddressConfig) mail.Address {
	if config == nil {
		config = &EmailAddressConfig{}
	}
	if len(config.EmailProviderDomainNameList) == 0 {
		config.EmailProviderDomainNameList = sample.EmailProviderDomainNames
	}

	firstName := FirstName(config.FirstNameList)
	lastName := LastName(config.LastNameList)
	emailProviderDomain := random.FromStringSlice(config.EmailProviderDomainNameList)
	address := firstName + "_" + lastName + "@" + emailProviderDomain
	address = strings.ReplaceAll(address, " ", "")
	address = strings.ToLower(address)
	return mail.Address{
		Name:    firstName + " " + lastName,
		Address: address,
	}
}
