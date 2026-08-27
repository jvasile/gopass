package action

import (
	"context"
	"time"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/pkg/debug"
)

// History displays the history of a given secret.
func (s *searchHandler) History(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)
	name := cmd.Args().Get(0)
	showPassword := cmd.Bool("password")

	if name == "" {
		return exit.Error(exit.Usage, nil, "Usage: %s history <NAME>", s.Name)
	}

	if !s.Store.Exists(ctx, name) {
		return exit.Error(exit.NotFound, nil, "Secret not found")
	}

	revs, err := s.Store.ListRevisions(ctx, name)
	if err != nil {
		return exit.Error(exit.Unknown, err, "Failed to get revisions: %s", err)
	}

	for _, rev := range revs {
		pw := ""
		if showPassword {
			_, sec, err := s.Store.GetRevision(ctx, name, rev.Hash)
			if err != nil {
				debug.Log("Failed to get revision %q of %q: %s", rev.Hash, name, err)
			}
			if err == nil {
				pw = " - " + sec.Password()
			}
		}
		out.Printf(ctx, "%s - %s <%s> - %s - %s%s\n", rev.Hash, rev.AuthorName, rev.AuthorEmail, rev.Date.Format(time.RFC3339), rev.Subject, pw)
	}

	return nil
}
