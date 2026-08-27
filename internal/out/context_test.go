package out

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jamesvasile.com/go/gopass/v2/internal/config"
)

func TestPrefix(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	assert.Empty(t, Prefix(ctx))

	ctx = AddPrefix(ctx, "[foo] ")
	assert.Equal(t, "[foo] ", Prefix(ctx))

	ctx = AddPrefix(ctx, "[bar] ")
	assert.Equal(t, "[foo] [bar] ", Prefix(ctx))

	ctx = AddPrefix(ctx, "")
	assert.Equal(t, "[foo] [bar] ", Prefix(ctx))
}

func TestNewline(t *testing.T) {
	t.Parallel()

	ctx := config.NewContextInMemory()

	assert.True(t, HasNewline(ctx))
	assert.False(t, HasNewline(WithNewline(ctx, false)))
}
