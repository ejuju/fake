package uuid

import (
	"github.com/google/uuid"
)

func NewUUIDWithFallback() string {
	return uuid.New().String()
}
