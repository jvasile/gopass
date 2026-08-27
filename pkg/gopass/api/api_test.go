package api_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/pkg/gopass/api"
	"jamesvasile.com/go/gopass/v2/pkg/gopass/secrets"
)

func Example() { //nolint:testableexamples
	ctx := config.NewContextInMemory()

	gp, err := api.New(ctx)
	if err != nil {
		panic(err)
	}

	// Listing secrets
	ls, err := gp.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, s := range ls {
		fmt.Printf("Secret: %s", s)
	}

	// Writing secrets
	sec := secrets.New()
	sec.SetPassword("foobar")
	if err := gp.Set(ctx, "my/new/secret", sec); err != nil {
		panic(err)
	}

	// Reading secrets
	sec, err = gp.Get(ctx, "my/new/secret", "latest")
	if err != nil {
		panic(err)
	}
	fmt.Printf("content of %s: %s\n", "my/new/secret", string(sec.Bytes()))

	// Removing a secret
	if err := gp.Remove(ctx, "my/new/secret"); err != nil {
		panic(err)
	}

	// Cleaning up
	if err := gp.Close(ctx); err != nil {
		panic(err)
	}
}

func TestApi(t *testing.T) {
	td := t.TempDir()
	t.Setenv("GOPASS_HOMEDIR", td)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithInteractive(ctx, false)

	gp, err := api.New(ctx)
	require.ErrorIs(t, err, api.ErrNotInitialized)
	assert.Nil(t, gp)

	// TODO: initialize store
}
