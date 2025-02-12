package constants

import (
	"os"
	"path/filepath"
)

const Name = "clicord"

const TmpFilePattern = Name + "_*.md"

var ConfigDirPath string

func init() {
	path, err := os.UserConfigDir()
	if err != nil {
		path = "."
	}
	ConfigDirPath = filepath.Join(path, Name)
}
