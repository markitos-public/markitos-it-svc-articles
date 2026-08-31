package types_test

import (
	"markitos-it-svc-faqs/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateValidTags(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	validTags, err := types.NewTags(tags)

	require.NoError(t, err)
	require.NotEmpty(t, validTags)
	require.NotNil(t, validTags)
	require.Equal(t, tags, validTags.Value())
}

func TestCantCreateEmptyTags(t *testing.T) {
	emptyTags, err := types.NewTags([]string{})

	require.Error(t, err)
	require.Nil(t, emptyTags)
}
