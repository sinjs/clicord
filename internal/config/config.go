package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sinjs/clicord/internal/constants"
)

type Config struct {
	Mouse bool `toml:"mouse"`

	Timestamps             bool   `toml:"timestamps"`
	TimestampsBeforeAuthor bool   `toml:"timestamps_before_author"`
	TimestampsFormat       string `toml:"timestamps_format"`

	UserAgent string `toml:"user_agent"`
	OS        string `toml:"os"`
	Browser   string `toml:"browser"`
	Device    string `toml:"device"`

	MessagesLimit uint8 `toml:"messages_limit"`

	Editor string `toml:"editor"`

	Keys  Keys  `toml:"keys"`
	Theme Theme `toml:"theme"`
}

func DefaultConfig() Config {
	return Config{
		Mouse: true,

		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.3",
		OS:        "Windows",
		Browser:   "Chrome",
		Device:    "",

		Timestamps:             true,
		TimestampsBeforeAuthor: true,
		TimestampsFormat:       time.Kitchen,

		MessagesLimit: 50,
		Editor:        "default",

		Keys:  defaultKeys(),
		Theme: defaultTheme(),
	}
}

// Reads the configuration file and parses it.
func Load() (*Config, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	cfg := DefaultConfig()
	path = filepath.Join(path, constants.Name, "config.toml")
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return &cfg, nil
	}

	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := toml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
