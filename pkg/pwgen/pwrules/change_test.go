package pwrules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jamesvasile.com/go/gopass/v2/internal/config"
)

func TestLookupChangeURL(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()
	assert.Equal(t, "https://account.gmx.net/ciss/security/edit/passwordChange", LookupChangeURL(ctx, "gmx.net"))
}
