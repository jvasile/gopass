package legacy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/crypto"
	_ "jamesvasile.com/go/gopass/v2/internal/backend/storage"
	"jamesvasile.com/go/gopass/v2/internal/config/legacy"
	"jamesvasile.com/go/gopass/v2/tests/gptest"
)

func TestNewConfig(t *testing.T) {
	u := gptest.NewUnitTester(t)
	assert.NotNil(t, u)

	t.Setenv("GOPASS_CONFIG", filepath.Join(os.TempDir(), ".gopass.yml"))

	cfg := legacy.New()
	cs := cfg.String()
	assert.Contains(t, cs, `&legacy.Config{AutoClip:false, AutoImport:false, ClipTimeout:45, ExportKeys:true, NoPager:false, Notifications:true,`)
	assert.Contains(t, cs, `SafeContent:false, Mounts:map[string]string{},`)

	cfg = &legacy.Config{
		Mounts: map[string]string{
			"foo": "",
			"bar": "",
		},
	}
	cs = cfg.String()
	assert.Contains(t, cs, `&legacy.Config{AutoClip:false, AutoImport:false, ClipTimeout:0, ExportKeys:false, NoPager:false, Notifications:false,`)
	assert.Contains(t, cs, `SafeContent:false, Mounts:map[string]string{"bar":"", "foo":""},`)
}

func TestSetConfigValue(t *testing.T) {
	u := gptest.NewUnitTester(t)
	assert.NotNil(t, u)

	t.Setenv("GOPASS_CONFIG", filepath.Join(os.TempDir(), ".gopass.yml"))

	cfg := legacy.New()
	require.NoError(t, cfg.SetConfigValue("autoclip", "true"))
	require.NoError(t, cfg.SetConfigValue("cliptimeout", "900"))
	require.NoError(t, cfg.SetConfigValue("path", "/tmp"))
	require.Error(t, cfg.SetConfigValue("autoclip", "yo"))
}
