package uuid

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUIDv4() string {
	return uuid.New().String()
}

func GenerateUUIDv4WithNoMinus() string {
	id := GenerateUUIDv4()
	return strings.ReplaceAll(id, "-", "")
}
