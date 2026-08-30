package types

import (
	"crypto/rand"
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidID = errors.New("invalid id must be an UUID v4 string")

type ID struct {
	value string
}

func NewID() (*ID, error) {
	return &ID{
		value: createUUID(),
	}, nil
}

func (id *ID) Value() string {
	return id.value
}

func NewIDFromString(s string) (*ID, error) {
	if s == "" {
		return nil, ErrInvalidID
	}

	id := ID{
		value: s,
	}

	if !id.IsValid() {
		return nil, ErrInvalidID
	}

	return &id, nil
}

func (id *ID) IsValid() bool {
	if id.value == "" {
		return false
	}

	return regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`).MatchString(id.value)
}

func createUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
