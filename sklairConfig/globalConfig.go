package sklairConfig

import (
	"os"
	"path/filepath"
)

// TODO: check TODO.md for more info about this

type GlobalConfig struct {
	//AutoUpdate bool `json:"autoUpdate,omitempty"`
}

var defaultGlobalConfig = GlobalConfig{
	//AutoUpdate: true,
}

func GlobalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".sklair/config.json"), nil
}
