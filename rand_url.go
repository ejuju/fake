package fake

import (
	"net/url"

	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

func URL(hosts []string, pagePaths []string) *url.URL {
	if len(hosts) == 0 {
		hosts = sample.PopularDomainNames
	}
	if len(pagePaths) == 0 {
		pagePaths = sample.PagePaths
	}

	return &url.URL{
		Scheme: "https",
		Host:   random.FromStringSlice(hosts),
		Path:   random.FromStringSlice(sample.PagePaths),
	}
}
