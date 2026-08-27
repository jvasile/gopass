//go:build windows

package termio

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
)

// promptPass will prompt user's for a password by terminal.
func promptPass(ctx context.Context, prompt string) (string, error) {
	if !ctxutil.IsTerminal(ctx) {
		return AskForString(ctx, prompt, "")
	}

	fmt.Fprintf(Stderr, "%s: ", prompt)
	passBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(Stderr, "")
	return string(passBytes), err
}
