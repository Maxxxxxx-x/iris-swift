package utils

import (
	"os"
)

func EnsureFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	return err
}
