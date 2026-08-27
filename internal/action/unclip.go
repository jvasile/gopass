package action

import (
	"context"
	"os"
	"time"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/pkg/clipboard"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// Unclip tries to erase the content of the clipboard.
func (s *miscHandler) Unclip(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)
	force := cmd.Bool("force")
	timeout := cmd.Int("timeout")
	name := os.Getenv("GOPASS_UNCLIP_NAME")
	checksum := os.Getenv("GOPASS_UNCLIP_CHECKSUM")

	time.Sleep(time.Second * time.Duration(timeout))

	mp := s.Store.MountPoint(name)
	ctx = config.WithMount(ctx, mp)

	if err := clipboard.Clear(ctx, name, checksum, force); err != nil {
		return exit.Error(exit.IO, err, "Failed to clear clipboard: %s", err)
	}

	return nil
}
