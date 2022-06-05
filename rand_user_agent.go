package fake

import (
	"github.com/ejuju/fake/internal/random"
)

func RandUserAgent(sampleUserAgents []string) string {
	if len(sampleUserAgents) == 0 {
		sampleUserAgents = SampleUserAgents
	}
	return random.FromStringSlice(sampleUserAgents)
}
