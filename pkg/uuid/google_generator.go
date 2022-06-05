package uuid

import (
	"github.com/google/uuid"
)

// GoogleUUIDGenerator implements the Generator interface
type GoogleUUIDGenerator struct{}

func NewGoogleUUIDGenerator() *GoogleUUIDGenerator {
	return &GoogleUUIDGenerator{}
}

func (gg *GoogleUUIDGenerator) Generate() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return "4631c17f-5f66-42f7-bb90-39b2619b8b44" // fallback uuid
	}
	return id.String()
}
