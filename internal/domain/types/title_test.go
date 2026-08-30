package types_test

import (
	"markitos-it-svc-articles/internal/domain/types"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateValidTitle(t *testing.T) {
	validTitle := "Valid Title"
	title, err := types.NewTitle(validTitle)

	require.NoError(t, err)
	require.NotEmpty(t, title)
	require.NotNil(t, title)
	require.Equal(t, validTitle, title.Value())
}

func TestCantCreateEmptyTitle(t *testing.T) {
	_, err := types.NewTitle("")

	require.Error(t, err)
}

func TestTitleHaveMin5CharsMax100Chars(t *testing.T) {
	invalidShortTitle := strings.Repeat("a", 4)
	shortTitle, err := types.NewTitle(invalidShortTitle)
	require.Error(t, err)
	require.Nil(t, shortTitle)

	invalidLongTitle := strings.Repeat("a", 101)
	longTitle, err := types.NewTitle(invalidLongTitle)
	require.Error(t, err)
	require.Nil(t, longTitle)

	validTitle := strings.Repeat("a", 10)
	title, err := types.NewTitle(validTitle)
	require.NoError(t, err)
	require.NotNil(t, title)
	require.Equal(t, validTitle, title.Value())
}
