package fake

import (
	"net/url"

	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

func RandURL(hosts []string, pageURIs []string) url.URL {
	if len(hosts) == 0 {
		hosts = sample.PopularDomainNames
	}
	if len(pageURIs) == 0 {
		pageURIs = sample.PageURIs
	}

	return url.URL{
		Scheme: "https",
		Host:   random.FromStringSlice(hosts),
		Path:   "/" + random.FromStringSlice(sample.PageURIs),
	}
}
