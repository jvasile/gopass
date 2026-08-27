package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/gopasspw/clipboard"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestCreate(t *testing.T) {
	u := gptest.NewUnitTester(t)

	clipboard.ForceUnsupported = true

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	require.NoError(t, act.cfg.Set("", "core.notifications", "false"))
	require.NoError(t, act.cfg.Set("", "core.cliptimeout", "1"))

	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	// create
	c := gptest.CliCtx(ctx, t)

	require.Error(t, act.Create(ctx, c))
	buf.Reset()
}
