package leaf

import (
	"bytes"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/backend/crypto/plain"
	"jamesvasile.com/go/gopass/v2/internal/backend/storage/fs"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/internal/recipients"
	"jamesvasile.com/go/gopass/v2/internal/store"
	"jamesvasile.com/go/gopass/v2/pkg/gopass/secrets"
)

func TestFsck(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	obuf := &bytes.Buffer{}
	out.Stdout = obuf
	defer func() {
		out.Stdout = os.Stdout
	}()

	// common setup
	tempdir := t.TempDir()

	s := &Store{
		alias:   "",
		path:    tempdir,
		crypto:  plain.New(),
		storage: fs.New(tempdir),
	}

	rs := recipients.New()
	rs.Add("john.doe")

	require.NoError(t, s.saveRecipients(ctx, rs, "test"))

	for _, e := range []string{"foo/bar", "foo/baz", "foo/zab"} {
		sec := secrets.NewAKV()
		sec.SetPassword("bar")
		require.NoError(t, s.Set(ctx, e, sec))
	}

	require.NoError(t, s.Fsck(ctx, "", nil))
	obuf.Reset()
}

func TestFsckCheckCaseConflicts(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	tempdir := t.TempDir()

	s := &Store{
		alias:   "",
		path:    tempdir,
		crypto:  plain.New(),
		storage: fs.New(tempdir),
	}

	rs := recipients.New()
	rs.Add("john.doe")
	require.NoError(t, s.saveRecipients(ctx, rs, "test"))

	// Set up entries with no case conflicts — should be fine.
	for _, e := range []string{"foo/bar", "foo/baz"} {
		sec := secrets.NewAKV()
		sec.SetPassword("x")
		require.NoError(t, s.Set(ctx, e, sec))
	}

	assert.NoError(t, s.fsckCheckCaseConflicts(ctx),
		"no case conflicts expected")

	// Now add entries that differ only in case.
	for _, e := range []string{"foo/Bar", "Foo/baz"} {
		sec := secrets.NewAKV()
		sec.SetPassword("x")
		if runtime.GOOS == "linux" {
			require.NoError(t, s.Set(ctx, e, sec), "Linux should allow case conflicts")
		} else {
			require.ErrorIs(t, s.Set(ctx, e, sec), store.ErrMeaninglessWrite)
		}
	}

	err := s.fsckCheckCaseConflicts(ctx)
	if runtime.GOOS == "linux" {
		// On case-sensitive filesystems the conflicting entries were stored
		// successfully, so the check must report them.
		assert.Error(t, err, "case conflicts should be reported")
	} else {
		// On case-insensitive filesystems (macOS, Windows) the Set calls
		// above returned ErrMeaninglessWrite, so no conflicting entries were
		// ever written and the check should find nothing.
		assert.NoError(t, err, "no case conflicts expected on case-insensitive filesystem")
	}
}

func TestCompareStringSlices(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		from    []string
		to      []string
		missing []string
		extra   []string
	}{
		{
			name:    "Add foo, remove baz",
			from:    []string{"foo", "bar"},
			to:      []string{"baz", "bar"},
			missing: []string{"foo"},
			extra:   []string{"baz"},
		},
		{
			name:    "Add foo, bar, baz, zab",
			from:    []string{"foo", "bar"},
			to:      []string{"foo", "bar", "bar", "baz", "zab"},
			missing: []string{},
			extra:   []string{"baz", "zab"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			missing, extra := compareStringSlices(tc.from, tc.to)
			assert.Equal(t, tc.missing, missing)
			assert.Equal(t, tc.extra, extra)
		})
	}
}
