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

func TestEdit(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithTerminal(ctx, false)
	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	// edit
	require.Error(t, act.Edit(ctx, gptest.CliCtx(ctx, t)))
	buf.Reset()

	// edit foo (existing)
	require.Error(t, act.Edit(ctx, gptest.CliCtx(ctx, t, "foo")))
	buf.Reset()

	// edit bar (new)
	require.Error(t, act.Edit(ctx, gptest.CliCtx(ctx, t, "foo")))
	buf.Reset()
}

func TestEditUpdate(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithTerminal(ctx, false)
	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	content := []byte("foobar")
	// no changes
	require.NoError(t, act.editUpdate(ctx, "foo", content, content, false, "test"))
	buf.Reset()

	// changes
	nContent := []byte("barfoo")
	require.NoError(t, act.editUpdate(ctx, "foo", content, nContent, false, "test"))
	buf.Reset()
}
