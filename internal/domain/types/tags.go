package types

import "errors"

var ErrInvalidTags = errors.New("invalid tags must be a non-empty slice of strings")

type Tags struct {
	value []string
}

func (t *Tags) Value() []string {
	return t.value
}

func NewTags(tags []string) (*Tags, error) {
	if len(tags) == 0 {
		return nil, ErrInvalidTags
	}

	return &Tags{
		value: tags,
	}, nil
}
