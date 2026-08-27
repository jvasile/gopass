package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/crypto"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/storage"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestUnclip(t *testing.T) {
	u := gptest.NewUnitTester(t)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	stdout = buf
	defer func() {
		out.Stdout = os.Stdout
		stdout = os.Stdout
	}()

	ctx := config.NewContextInMemory()
	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	t.Run("unlcip should fail", func(t *testing.T) {
		require.Error(t, act.Unclip(ctx, gptest.CliCtxWithFlags(ctx, t, map[string]string{"timeout": "0"})))
	})
}
