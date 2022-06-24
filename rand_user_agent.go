package fake

import (
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/internal/sample"
)

func RandUserAgent(sampleUserAgents []string) string {
	if len(sampleUserAgents) == 0 {
		sampleUserAgents = sample.UserAgents
	}
	return random.FromStringSlice(sampleUserAgents)
}
