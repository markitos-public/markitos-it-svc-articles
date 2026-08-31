package types_test

import (
	"markitos-it-svc-faqs/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGenerateID(t *testing.T) {
	sut, err := types.NewID()

	require.NoError(t, err)
	require.NotEmpty(t, sut)
	require.NotNil(t, sut)
	require.True(t, sut.IsValid())
}

func TestCanValidateID(t *testing.T) {
	id := "a4f1f7f2-dc93-4e67-8d15-1212b98c8736"

	validID, errValid := types.NewIDFromString(id)
	require.NoError(t, errValid)
	require.True(t, validID.IsValid())

	invalidID, errInvalid := types.NewIDFromString("invalid-id")
	require.Error(t, errInvalid)
	require.Nil(t, invalidID)
	require.Equal(t, validID.Value(), id)
}

func TestCantCreateInvalidID(t *testing.T) {
	invalidID, err := types.NewIDFromString("invalid-id")

	require.Error(t, err)
	require.Equal(t, types.ErrInvalidID, err)
	require.Nil(t, invalidID)
}
