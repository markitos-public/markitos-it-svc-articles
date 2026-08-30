package types_test

import (
	"markitos-it-svc-articles/internal/domain/types"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateContent(t *testing.T) {
	validContent := "This is a valid content string."
	content, err := types.NewContent(validContent)

	require.NoError(t, err)
	require.Equal(t, validContent, content.Value())
}

func TestCannotCreateContentWithInvalidLength(t *testing.T) {
	shortContent, err := types.NewContent(strings.Repeat("a", 9))

	require.Error(t, err)
	require.Equal(t, types.ErrInvalidContent, err)
	require.Nil(t, shortContent)
}
