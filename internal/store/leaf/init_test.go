package leaf

import (
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
)

func TestInit(t *testing.T) {
	ctx := config.NewContextInMemory()

	s, err := createSubStore(t)
	require.NoError(t, err)
	require.Error(t, s.Init(ctx, "", "0xDEADBEEF"))
}
