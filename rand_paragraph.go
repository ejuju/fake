package fake

import (
	"github.com/ejuju/fake/internal/random"
)

func RandParagraph(sampleParagraphs []string) string {
	if len(sampleParagraphs) == 0 {
		sampleParagraphs = SampleEnglishSentences
	}
	return random.FromStringSlice(sampleParagraphs)
}
