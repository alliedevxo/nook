package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetUserConfigDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("user config dir: %w", err)
	}

	dir := filepath.Join(base, "nook")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("making app dir: %w", err)
	}

	return filepath.Join(dir, "nook.db"), nil
}