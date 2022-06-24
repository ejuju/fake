package fake

import (
	"net/mail"
	"strings"

	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

type RandEmailAddressConfig struct {
	FirstNameList               []string
	LastNameList                []string
	EmailProviderDomainNameList []string
}

func RandEmailAddress(config *RandEmailAddressConfig) mail.Address {
	if config == nil {
		config = &RandEmailAddressConfig{}
	}
	if len(config.EmailProviderDomainNameList) == 0 {
		config.EmailProviderDomainNameList = sample.EmailProviderDomainNames
	}

	firstName := RandFirstName(config.FirstNameList)
	lastName := RandLastName(config.FirstNameList)
	emailProviderDomain := random.FromStringSlice(config.EmailProviderDomainNameList)
	address := firstName + "_" + lastName + "@" + emailProviderDomain
	address = strings.ReplaceAll(address, " ", "")
	address = strings.ToLower(address)
	return mail.Address{
		Name:    firstName + " " + lastName,
		Address: address,
	}
}
