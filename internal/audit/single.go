package audit

import (
	"context"
	"fmt"

	"github.com/muesli/crunchy"
	"jamesvasile.com/go/gopass/v2/internal/out"
)

// Single runs a password strength audit on a single password.
func Single(ctx context.Context, password string) {
	validator := crunchy.NewValidator()
	if err := validator.Check(password); err != nil {
		out.Printf(ctx, fmt.Sprintf("Warning: %s", err))
	}
}
