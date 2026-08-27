package termio

import (
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

func TestPromptPass(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithTerminal(ctx, false)
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	_, err := promptPass(ctx, "foo")
	require.NoError(t, err)
}
