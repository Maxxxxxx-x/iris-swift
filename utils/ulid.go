package utils

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

func GenerateULID() (ulid.ULID, error) {
	entropy := rand.Reader
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}


