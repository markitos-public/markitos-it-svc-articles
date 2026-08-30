package types

import "errors"

var ErrInvalidContent = errors.New("invalid content must be a non-empty string between 10 and 1000 characters")

type Content struct {
	value string
}

func NewContent(s string) (*Content, error) {
	if len(s) < 10 || len(s) > 1000 {
		return nil, ErrInvalidContent
	}

	return &Content{
		value: s,
	}, nil
}

func (c *Content) Value() string {
	return c.value
}
