package leaf

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
)

func TestGPG(t *testing.T) {
	ctx := config.NewContextInMemory()

	obuf := &bytes.Buffer{}
	out.Stdout = obuf
	defer func() {
		out.Stdout = os.Stdout
	}()

	s, err := createSubStore(t)
	require.NoError(t, err)

	require.NoError(t, s.ImportMissingPublicKeys(ctx))

	newRecp := "A3683834"
	err = s.AddRecipient(ctx, newRecp)
	require.NoError(t, err)

	require.NoError(t, s.ImportMissingPublicKeys(ctx))
}
