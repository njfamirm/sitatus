package jsonLine

import (
	"os"
	"path/filepath"
)

func CreateFolder(path string) error {
	// check if folder already exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	// create parent directories recursively
	parent := filepath.Dir(path)
	if err := CreateFolder(parent); err != nil {
		return err
	}

	// create folder
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}

	return nil
}
