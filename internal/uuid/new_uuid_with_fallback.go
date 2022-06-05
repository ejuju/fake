package uuid

import (
	"github.com/google/uuid"
)

func NewUUIDWithFallback() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return "4631c17f-5f66-42f7-bb90-39b2619b8b44" // fallback uuid
	}
	return id.String()
}
