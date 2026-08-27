package gitfs

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

func TestGitConfig(t *testing.T) {
	gitdir := filepath.Join(t.TempDir(), "git")
	require.NoError(t, os.Mkdir(gitdir, 0o755))

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	git, err := Init(ctx, gitdir, "Dead Beef", "dead.beef@example.org")
	require.NoError(t, err)
	un, err := git.ConfigGet(ctx, "user.name")
	require.NoError(t, err)
	assert.Equal(t, "Dead Beef", un)

	require.NoError(t, git.InitConfig(ctx, "Foo Bar", "foo.bar@example.org"))
	un, err = git.ConfigGet(ctx, "user.name")
	require.NoError(t, err)
	assert.Equal(t, "Foo Bar", un)

	require.NoError(t, git.ConfigSet(ctx, "user.name", "foo"))
	un, err = git.ConfigGet(ctx, "user.name")
	require.NoError(t, err)
	assert.Equal(t, "foo", un)
}
