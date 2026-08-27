package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestSync(t *testing.T) {
	u := gptest.NewUnitTester(t)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	out.Stderr = buf
	defer func() {
		out.Stdout = os.Stdout
		out.Stderr = os.Stderr
	}()

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	t.Run("default", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.Sync(ctx, gptest.CliCtx(ctx, t)))
	})

	t.Run("sync --store=root", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.Sync(ctx, gptest.CliCtxWithFlags(ctx, t, map[string]string{"store": "root"})))
	})
}
