package cli

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(config.NewContextInMemory())

	g := &GPG{}
	g.binary = "rundll32"

	_, err := g.Encrypt(ctx, []byte("foo"), nil)

	// No recipients are configured so it will fail
	require.Error(t, err)
	cancel()
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(config.NewContextInMemory())

	g := &GPG{}
	g.binary = "rundll32"

	_, err := g.Decrypt(ctx, []byte("foo"))
	require.NoError(t, err)
	cancel()
}

func TestGenerateIdentity(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(config.NewContextInMemory())

	g := &GPG{}
	g.binary = "rundll32"

	_, err := g.GenerateIdentity(ctx, "foo", "foo@bar.com", "bar")
	require.NoError(t, err)
	cancel()
}
