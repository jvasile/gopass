package action

import (
	"context"
	"sort"
	"strings"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/pwgen/pwrules"
)

// AliasesPrint prints all configured aliases for password generation rules.
func (s *miscHandler) AliasesPrint(ctx context.Context, cmd *cli.Command) error {
	out.Printf(ctx, "Configured aliases:")
	aliases := pwrules.AllAliases(ctx)
	keys := make([]string, 0, len(aliases))
	for k := range aliases {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		out.Printf(ctx, "- %s -> %s", k, strings.Join(aliases[k], ", "))
	}

	return nil
}
