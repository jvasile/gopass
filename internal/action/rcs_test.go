package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestGit(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithInteractive(ctx, false)

	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	stdout = buf
	defer func() {
		out.Stdout = os.Stdout
		stdout = os.Stdout
	}()

	// git init
	c := gptest.CliCtxWithFlags(ctx, t, map[string]string{"name": "foobar", "email": "foo.bar@example.org"})
	require.NoError(t, act.RCSInit(ctx, c))
	buf.Reset()

	// getUserData
	name, email := act.getUserData(ctx, "", "", "")
	assert.Equal(t, "0xDEADBEEF", name)
	assert.Equal(t, "0xDEADBEEF", email)

	// GitAddRemote
	require.Error(t, act.RCSAddRemote(ctx, c))
	buf.Reset()

	// GitRemoveRemote
	require.Error(t, act.RCSRemoveRemote(ctx, c))
	buf.Reset()

	// GitPull
	require.Error(t, act.RCSPull(ctx, c))
	buf.Reset()

	// GitPush
	require.NoError(t, act.RCSPush(ctx, c))
	buf.Reset()
}
