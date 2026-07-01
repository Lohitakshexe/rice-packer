package manifest

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Manifest struct {
	Name          string            `toml:"name"`
	Author        string            `toml:"author"`
	Version       string            `toml:"version"`
	DisplayServer string            `toml:"display_server"`
	WindowManager string            `toml:"window_manager"`
	Configs       []string          `toml:"configs"`
	Dependencies  []string          `toml:"dependencies"`
	Hardware      map[string]string `toml:"hardware"`
}

func Load(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m Manifest
	err = toml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
