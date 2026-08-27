package exit

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/out"
)

func TestError(t *testing.T) {
	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	require.Error(t, Error(Unknown, fmt.Errorf("test"), "test"))
	assert.NotContains(t, buf.String(), "Stacktrace")
}
