package age

import (
	"context"

	"filippo.io/age/plugin"
	"jamesvasile.com/go/gopass/v2/internal/cui"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/termio"
)

var pluginTerminalUI = &plugin.ClientUI{
	DisplayMessage: func(name, message string) error {
		out.Printf(context.Background(), "%s plugin: %s", name, message)

		return nil
	},
	RequestValue: func(name, message string, _ bool) (string, error) {
		var err error
		defer func() {
			if err != nil {
				out.Warningf(context.Background(), "could not read value for age-plugin-%s: %v", name, err)
			}
		}()
		secret, err := termio.AskForPassword(context.Background(), "secret", false)
		if err != nil {
			return "", err
		}

		return secret, nil
	},
	Confirm: func(name, message, yes, no string) (bool, error) {
		rep, _ := cui.GetSelection(context.Background(), message, []string{yes, no})
		if rep == yes {
			return true, nil
		}

		return false, nil
	},

	WaitTimer: func(name string) {
		out.Printf(context.Background(), "waiting on %s plugin...", name)
	},
}
