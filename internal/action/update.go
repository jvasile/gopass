package action

import (
	"context"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/internal/updater"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// Update will start the interactive update assistant.
func (s *miscHandler) Update(ctx context.Context, cmd *cli.Command) error {
	_ = s.rem.Reset("update")

	ctx = ctxutil.WithGlobalFlags(ctx, cmd)

	if s.version.String() == "0.0.0+HEAD" {
		out.Errorf(ctx, "Cannot check version against HEAD")

		return nil
	}

	out.Printf(ctx, "⚒ Checking for available updates ...")
	if err := updater.Update(ctx, s.version); err != nil {
		return exit.Error(exit.Unknown, err, "Failed to update gopass: %s", err)
	}

	out.OKf(ctx, "gopass is up to date")

	return nil
}
