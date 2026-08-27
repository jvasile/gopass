package action

import (
	"context"

	"github.com/urfave/cli/v3"
	"jamesvasile.com/go/gopass/v2/internal/action/exit"
	"jamesvasile.com/go/gopass/v2/internal/backend"
	"jamesvasile.com/go/gopass/v2/internal/backend/crypto/age"
	"jamesvasile.com/go/gopass/v2/internal/config"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/pkg/ctxutil"
	"jamesvasile.com/go/gopass/v2/pkg/debug"
	"jamesvasile.com/go/gopass/v2/pkg/termio"
)

// Convert converts a store to a different set of backends.
func (s *miscHandler) Convert(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithGlobalFlags(ctx, cmd)
	ctx = age.WithOnlyNative(ctx, true)

	store := cmd.String("store")
	move := cmd.Bool("move")

	sub, err := s.Store.GetSubStore(store)
	if err != nil {
		return exit.Error(exit.NotFound, err, "mount %q not found: %s", store, err)
	}

	// we know it's a valid mount at this point
	ctx = config.WithMount(ctx, store)

	oldStorage := sub.Storage().Name()

	storage, err := backend.StorageRegistry.Backend(oldStorage)
	if err != nil {
		return exit.Error(exit.Unknown, err, "unknown source storage backend %q: %s", oldStorage, err)
	}

	if sv := cmd.String("storage"); sv != "" {
		var err error
		storage, err = backend.StorageRegistry.Backend(sv)
		if err != nil {
			return exit.Error(exit.Usage, err, "unknown destination storage backend %q: %s", sv, err)
		}
	}

	oldCrypto := sub.Crypto().Name()

	crypto, err := backend.CryptoRegistry.Backend(oldCrypto)
	if err != nil {
		return exit.Error(exit.Unknown, err, "unknown source crypto backend %q: %s", oldCrypto, err)
	}

	if sv := cmd.String("crypto"); sv != "" {
		var err error
		crypto, err = backend.CryptoRegistry.Backend(sv)
		if err != nil {
			return exit.Error(exit.Usage, err, "unknown destination crypto backend %q: %s", sv, err)
		}
	}

	if oldCrypto == crypto.String() && oldStorage == storage.String() {
		out.Notice(ctx, "No conversion needed. Source and destination match.")

		return nil
	}

	if oldCrypto != crypto.String() {
		debug.Log("attempting to convert crypto from %q to %q", oldCrypto, crypto.String())

		cbe, err := backend.NewCrypto(ctx, crypto)
		if err != nil {
			return err
		}

		if err := s.initCheckPrivateKeysFn(ctx, cbe); err != nil {
			return err
		}
		out.Printf(ctx, "Crypto %q has private keys", crypto.String())
	}

	out.Noticef(ctx, "Converting %q. Crypto: %q -> %q, Storage: %q -> %q", store, oldCrypto, crypto, oldStorage, storage)
	ok, err := termio.AskForBool(ctx, "Continue?", false)
	if err != nil {
		return err
	}
	if ctxutil.IsInteractive(ctx) && !ok {
		out.Notice(ctx, "Aborted")

		return nil
	}

	if err := s.Store.Convert(ctx, store, crypto, storage, move); err != nil {
		return exit.Error(exit.Unknown, err, "failed to convert store %q: %s", store, err)
	}

	out.OKf(ctx, "Successfully converted %q", store)

	return nil
}
