package cui

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/backend/crypto/plain"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestAskForPrivateKey(t *testing.T) {
	buf := &bytes.Buffer{}
	Stdout = buf
	defer func() {
		Stdout = os.Stdout
	}()

	ctx := config.NewContextInMemory()

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	key, err := AskForPrivateKey(ctx, plain.New(), "test")
	require.NoError(t, err)
	assert.Equal(t, "0xDEADBEEF", key)
	buf.Reset()
}

func TestAskForGitConfigUser(t *testing.T) {
	// necessary for setting up the env
	u := gptest.NewGUnitTester(t)
	assert.NotNil(t, u)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithTerminal(ctx, true)
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	_, _, err := AskForGitConfigUser(ctx, plain.New())
	require.NoError(t, err)
}

type fakeMountPointer []string

func (f fakeMountPointer) MountPoints() []string {
	return f
}

func TestAskForStore(t *testing.T) {
	ctx := config.NewContextInMemory()

	// test non-interactive
	ctx = ctxutil.WithInteractive(ctx, false)
	assert.Empty(t, AskForStore(ctx, fakeMountPointer{"foo", "bar"}))

	// test interactive
	ctx = ctxutil.WithInteractive(ctx, true)
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	assert.Empty(t, AskForStore(ctx, fakeMountPointer{"foo", "bar"}))

	// test zero mps
	assert.Empty(t, AskForStore(ctx, fakeMountPointer{}))

	// test one mp
	assert.Empty(t, AskForStore(ctx, fakeMountPointer{"foo"}))

	// test two mps
	assert.Empty(t, AskForStore(ctx, fakeMountPointer{"foo", "bar"}))
}

func TestSorted(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []string{"a", "b", "c"}, sorted([]string{"c", "a", "b"}))
}
