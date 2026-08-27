package action

import (
	"context"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// Link creates a symlink.
func (s *secretHandler) Link(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)

	from := cmd.Args().Get(0)
	to := cmd.Args().Get(1)

	if from == "" || to == "" {
		return exit.Error(exit.Usage, nil, "Usage: link <from> <to>")
	}

	return s.Store.Link(ctx, from, to)
}
