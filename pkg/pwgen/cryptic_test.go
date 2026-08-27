package pwgen

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/pwgen/pwrules"
)

func TestCrypticForDomain(t *testing.T) {
	t.Parallel()

	rules := pwrules.AllRules()
	keys := make([]string, 0, len(rules))

	for k := range rules {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, domain := range keys {
		t.Run(domain, func(t *testing.T) {
			for _, length := range []int{1, 4, 8, 100} {
				tcName := fmt.Sprintf("%s: generated password with %d chars", domain, length)
				c := NewCrypticForDomain(config.NewContextInMemory(), length, domain)
				c.MaxTries = 1024

				require.NotNil(t, c, tcName)

				pw := c.Password()

				assert.NotEmpty(t, pw, tcName)
				t.Logf("%s -> %s (%d)", tcName, pw, len(pw))
			}
		})
	}
}

func TestUniqueChars(t *testing.T) {
	t.Parallel()

	for in, out := range map[string]string{
		"foobar": "abfor",
		"abced":  "abcde",
	} {
		assert.Equal(t, out, uniqueChars(in))
	}
}
