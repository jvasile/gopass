//go:build !windows

package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	g := &GPG{}
	g.binary = "true"

	_, err := g.Encrypt(ctx, []byte("foo"), nil)
	// No recipients are configured so it will fail
	require.Error(t, err)
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	g := &GPG{}
	g.binary = "true"

	_, err := g.Decrypt(ctx, []byte("foo"))
	require.NoError(t, err)
}

func TestGenerateIdentity(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	g := &GPG{}
	g.binary = "true"

	_, err := g.GenerateIdentity(ctx, "foo", "foo@bar.com", "bar")
	require.NoError(t, err)
}
