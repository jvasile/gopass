package secrets

import (
	"jamesvasile.com/go/gopass/v2/pkg/gopass"
)

// New creates a new secret.
// It returns a new AKV secret.
func New() gopass.Secret { //nolint:ireturn
	return NewAKV()
}
