package email

import (
	"math/rand"
	"net/mail"
	"strings"
)

type AddressGenerator struct {
	config AddressGeneratorConfig
}

type AddressGeneratorConfig struct {
	FirstNames               []string
	LastNames                []string
	EmailProviderDomainNames []string
}

func NewEmailAddressGenerator(config AddressGeneratorConfig) *AddressGenerator {
	return &AddressGenerator{config: config}
}

func (ag *AddressGenerator) Generate() mail.Address {
	firstName := ag.config.FirstNames[rand.Intn(len(ag.config.FirstNames))]
	lastName := ag.config.LastNames[rand.Intn(len(ag.config.LastNames))]
	emailProviderDomain := ag.config.EmailProviderDomainNames[rand.Intn(len(ag.config.EmailProviderDomainNames))]
	address := firstName + "_" + lastName + "@" + emailProviderDomain
	address = strings.ReplaceAll(address, " ", "")
	address = strings.ToLower(address)
	return mail.Address{
		Name:    firstName + " " + lastName,
		Address: address,
	}
}
