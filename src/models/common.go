package models

import (
	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}
