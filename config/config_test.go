package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"

	"github.com/adrg/xdg"
)

func TestSetGetValue(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	xdg.Reload()

	err := SetValue("team", "my-team")
	if err != nil {
		t.Fatal(err)
	}

	got, err := GetValue("team")
	if err != nil {
		t.Fatal(err)
	}
	if got != "my-team" {
		t.Fatalf("expected %q, got %q", "my-team", got)
	}

	err = UnsetValue("team")
	if err != nil {
		t.Fatal(err)
	}

	got, err = GetValue("team")
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("expected %q, got %q", "", got)
	}
}

func TestSetConfigHome(t *testing.T) {
	r := require.New(t)
	configDir := t.TempDir()

	err := SetConfigHome(configDir)
	r.NoError(err)

	r.Equal(configDir, xdg.ConfigHome)

	err = SetValue("team", "set-config-test-value")
	r.NoError(err)

	// check that the config file was created in the temporary directory,
	// not somewhere else
	_, err = os.Stat(path.Join(configDir, "loki", "config.json"))
	r.NoError(err)

	err = UnsetConfigHome()
	r.NoError(err)

	// check that we are no longer set to the temporary directory
	r.NotEqual(configDir, xdg.ConfigHome)
}

func TestSetDataHome(t *testing.T) {
	r := require.New(t)
	configDir := t.TempDir()

	err := SetDataHome(configDir)
	r.NoError(err)

	r.Equal(configDir, xdg.DataHome)

	err = SaveDataString("loki/token", "my-token")
	r.NoError(err)

	// check that the config file was created in the temporary directory,
	// not somewhere else
	_, err = os.Stat(path.Join(configDir, "loki", "token"))
	r.NoError(err)

	err = UnsetDataHome()
	r.NoError(err)

	// check that we are no longer set to the temporary directory
	r.NotEqual(configDir, xdg.DataHome)
}
