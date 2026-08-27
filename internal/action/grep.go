package action

import (
	"context"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/internal/tree"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// Grep searches a string inside the content of all files.
func (s *searchHandler) Grep(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)
	if !cmd.Args().Present() {
		return exit.Error(exit.Usage, nil, "Usage: %s grep arg", s.Name)
	}

	// get the search term.
	needle := cmd.Args().First()

	haystack, err := s.Store.List(ctx, tree.INF)
	if err != nil {
		return exit.Error(exit.List, err, "failed to list store: %s", err)
	}

	matchFn := func(haystack string) bool {
		return strings.Contains(haystack, needle)
	}

	if cmd.Bool("regexp") {
		re, err := regexp.Compile(needle)
		if err != nil {
			return exit.Error(exit.Usage, err, "failed to compile regexp %q: %s", needle, err)
		}
		matchFn = re.MatchString
	}

	var matches int
	var errors int
	for _, v := range haystack {
		sec, err := s.Store.Get(ctx, v)
		if err != nil {
			out.Errorf(ctx, "Failed to decrypt %s: %v", v, err)
			errors++

			continue
		}

		if matchFn(string(sec.Bytes())) {
			out.Printf(ctx, "%s matches", color.BlueString(v))
			matches++
		}
	}

	if errors > 0 {
		out.Warningf(ctx, "%d secrets failed to decrypt", errors)
	}
	out.Printf(ctx, "\nScanned %d secrets. %d matches, %d errors", len(haystack), matches, errors)

	return nil
}
