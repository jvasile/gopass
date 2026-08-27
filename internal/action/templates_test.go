package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/crypto"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/storage"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestTemplates(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := config.NewContextInMemory()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithTerminal(ctx, false)

	act, err := newMock(ctx, u.StoreDir(""))
	require.NoError(t, err)
	require.NotNil(t, act)
	ctx = act.cfg.WithConfig(ctx)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	stdout = buf
	color.NoColor = true
	defer func() {
		stdout = os.Stdout
		out.Stdout = os.Stdout
	}()

	t.Run("display empty template tree", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.TemplatesPrint(ctx, gptest.CliCtx(ctx, t, "foo")))
		assert.Equal(t, "gopass\n\n", buf.String())
	})

	t.Run("add template", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.Store.SetTemplate(ctx, "foo", []byte("foobar")))
		require.NoError(t, act.TemplatesPrint(ctx, gptest.CliCtx(ctx, t, "foo")))
		want := `gopass
└── foo

`
		assert.Contains(t, buf.String(), want)
	})

	t.Run("complete templates", func(t *testing.T) {
		defer buf.Reset()
		act.TemplatesComplete(ctx, gptest.CliCtx(ctx, t, "foo"))
		assert.Equal(t, "foo\n", buf.String())
	})

	t.Run("print template", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.TemplatePrint(ctx, gptest.CliCtx(ctx, t, "foo")))
		assert.Equal(t, "foobar\n", buf.String())
	})

	t.Run("edit template", func(t *testing.T) {
		defer buf.Reset()
		require.Error(t, act.TemplateEdit(ctx, gptest.CliCtx(ctx, t, "foo")))
	})

	t.Run("remove template", func(t *testing.T) {
		defer buf.Reset()
		require.NoError(t, act.TemplateRemove(ctx, gptest.CliCtx(ctx, t, "foo")))
	})
}
