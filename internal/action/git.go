package action

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// Git passes the git command to the underlying backend.
func (s *syncHandler) Git(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)
	store := cmd.String("store")

	sub, err := s.Store.GetSubStore(store)
	if err != nil || sub == nil {
		return exit.Error(exit.Git, err, "failed to get sub store %s: %s", store, err)
	}

	args := cmd.Args().Slice()
	out.Noticef(ctx, "Running 'git %s' in %s...", strings.Join(args, " "), sub.Path())
	gitCmd := exec.CommandContext(ctx, "git", args...)
	gitCmd.Dir = sub.Path()
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	gitCmd.Stdin = os.Stdin

	return gitCmd.Run()
}
