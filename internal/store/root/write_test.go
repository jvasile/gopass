package root

import (
	"testing"

	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/pkg/gopass/secrets"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestSet(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)

	rs, err := createRootStore(ctx, u)
	require.NoError(t, err)

	sec := secrets.NewAKV()
	sec.SetPassword("foo")
	_, err = sec.Write([]byte("bar"))
	require.NoError(t, err)
	require.NoError(t, rs.Set(ctx, "zab", sec))

	err = rs.Set(ctx, "zab2", sec)
	require.NoError(t, err)
}
