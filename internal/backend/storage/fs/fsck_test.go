package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

func TestFsck(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithHidden(ctx, true)

	path := t.TempDir()

	l := &loader{}
	s, err := l.Init(ctx, path)
	require.NoError(t, err)
	require.NoError(t, l.Handles(ctx, path))

	for _, fn := range []string{
		filepath.Join(path, ".plain-ids"),
		filepath.Join(path, "foo", "bar"),
		filepath.Join(path, "foo", "zen"),
	} {
		require.NoError(t, os.MkdirAll(filepath.Dir(fn), 0o777))
		require.NoError(t, os.WriteFile(fn, []byte(fn), 0o663))
	}

	require.NoError(t, s.Fsck(ctx))
}
