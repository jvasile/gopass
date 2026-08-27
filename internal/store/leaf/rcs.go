package leaf

import (
	"context"
	"fmt"

	"jamesvasile.com/go/gopass/v2/internal/backend"
	"jamesvasile.com/go/gopass/v2/internal/out"
	"jamesvasile.com/go/gopass/v2/internal/store"
	"jamesvasile.com/go/gopass/v2/pkg/debug"
	"jamesvasile.com/go/gopass/v2/pkg/gopass"
	"jamesvasile.com/go/gopass/v2/pkg/gopass/secrets/secparse"
)

// GitInit initializes the git storage.
func (s *Store) GitInit(ctx context.Context) error {
	storage, err := backend.InitStorage(ctx, backend.GetStorageBackend(ctx), s.path)
	if err != nil {
		return err
	}
	s.storage = storage

	return nil
}

// ListRevisions will list all revisions for a secret.
func (s *Store) ListRevisions(ctx context.Context, name string) ([]backend.Revision, error) {
	p := s.passfile(ctx, name)

	return s.storage.Revisions(ctx, p)
}

// GetRevision will retrieve a single revision from the backend.
func (s *Store) GetRevision(ctx context.Context, name, revision string) (gopass.Secret, error) {
	p := s.passfile(ctx, name)
	ciphertext, err := s.storage.GetRevision(ctx, p, revision)
	if err != nil {
		return nil, fmt.Errorf("failed to get ciphertext of %q@%q: %w", name, revision, err)
	}

	content, err := s.crypto.Decrypt(ctx, ciphertext)
	if err != nil {
		debug.Log("Decryption failed: %s", err)

		return nil, store.ErrDecrypt
	}

	sec, err := secparse.Parse(content)
	if err != nil {
		debug.Log("Failed to parse YAML: %s", err)
	}

	return sec, nil
}

// GitStatus shows the git status output.
func (s *Store) GitStatus(ctx context.Context, _ string) error {
	debug.Log("RCS status for %s", s.path)
	buf, err := s.storage.Status(ctx)
	if err != nil {
		debug.Log("RCS status failed for %s: %s", s.path, err)

		return fmt.Errorf("failed to get RCS status for %s: %w", s.path, err)
	}

	out.Printf(ctx, string(buf))

	return nil
}
