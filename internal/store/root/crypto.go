package root

import (
	"context"

	"jamesvasile.com/go/gopass/v2/internal/backend"
	"jamesvasile.com/go/gopass/v2/pkg/debug"
)

// Crypto returns the crypto backend.
func (r *Store) Crypto(ctx context.Context, name string) backend.Crypto {
	sub, _ := r.getStore(name)
	if !sub.Valid() {
		debug.Log("Sub-Store not found for %s. Returning nil crypto backend", name)

		return nil
	}

	return sub.Crypto()
}
