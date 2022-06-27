package fake

import (
	"github.com/ejuju/fake/internal/random"
	"github.com/ejuju/fake/pkg/sample"
)

func Paragraph(sampleParagraphs []string) string {
	if len(sampleParagraphs) == 0 {
		sampleParagraphs = sample.EnglishSentences
	}
	return random.FromStringSlice(sampleParagraphs)
}
