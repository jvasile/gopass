package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/crypto"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/storage"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestVersion(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithInteractive(ctx, false)

	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	stdout = buf
	defer func() {
		out.Stdout = os.Stdout
		stdout = os.Stdout
	}()

	cli.VersionPrinter = func(*cli.Command) {
		out.Printf(ctx, "gopass version 0.0.0-test")
	}

	t.Run("print fixed version", func(t *testing.T) {
		require.NoError(t, act.Version(ctx, gptest.CliCtx(ctx, t)))
	})
}
