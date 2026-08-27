package clipboard

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

func TestNotExistingClipboardClearCommand(t *testing.T) {
	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	t.Setenv("GOPASS_CLIPBOARD_CLEAR_CMD", "not_existing_command")

	maybeErr := Clear(ctx, "", "", false)
	require.Error(t, maybeErr)
	assert.Contains(t, maybeErr.Error(), "\"not_existing_command\": executable file not found in")
}

func TestUnclip(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	buf := &bytes.Buffer{}
	out.Stdout = buf

	defer func() {
		out.Stdout = os.Stdout
	}()

	require.EqualError(t, Clear(ctx, "", "", false), ErrNotSupported.Error())
}
