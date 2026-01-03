package sklairConfig

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// TODO: expand this struct when JS obfuscation is added
type ObfuscateJS struct {
	Enabled bool `json:"enabled,omitempty"`
}

type PreventFOUC struct {
	Enabled bool   `json:"enabled,omitempty"`
	Colour  string `json:"colour,omitempty"`
}

//type ResourceHints struct {
//	Enabled    bool   `json:"enabled,omitempty"`
//	SiteOrigin string `json:"siteOrigin,omitempty"`
//}

type ProjectConfig struct {
	Hooks string `json:"hooks,omitempty"`

	Input      string `json:"input,omitempty"`
	Components string `json:"components,omitempty"`

	Exclude        []string `json:"exclude,omitempty"`
	ExcludeCompile []string `json:"excludeCompile,omitempty"`

	Output string `json:"output,omitempty"`

	Minify      bool         `json:"minify,omitempty"`
	ObfuscateJS *ObfuscateJS `json:"obfuscateJS,omitempty"`

	PreventFOUC *PreventFOUC `json:"PreventFOUC,omitempty"`
	//ResourceHints *ResourceHints `json:"resourceHints,omitempty"` // TODO: in sklair init, add ResourceHints to the questionnaire
}

var DefaultConfig = ProjectConfig{
	Hooks: "", // disabled by default

	Input:      "./",
	Components: "./components",

	Exclude:        []string{},
	ExcludeCompile: []string{},

	Output: "./build",

	Minify: false,
	ObfuscateJS: &ObfuscateJS{
		Enabled: false,
	},

	PreventFOUC: &PreventFOUC{
		Enabled: false,
		Colour:  "#202020",
	},
	//ResourceHints: &ResourceHints{
	//	Enabled:    false,
	//	SiteOrigin: "https://sklair.numelon.com", // TODO: maybe just make it empty by default
	//},
}

func resolveProjectConfigPath() string {
	if _, err := os.Stat("sklair.json"); err == nil {
		return "sklair.json"
	}

	if _, err := os.Stat("src/sklair.json"); err == nil {
		return "src/sklair.json"
	}

	return "sklair.json" // default
}

func LoadProjectConfig() (*ProjectConfig, string, error) {
	configPath := resolveProjectConfigPath()

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, "", err
	}

	config := DefaultConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, "", err
	}

	return &config, filepath.Dir(configPath), nil
}
