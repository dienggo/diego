package helper

import (
	"errors"
	"os"
)

// CheckFileExist to check file exist or not exist
func CheckFileExist(filePath string) (error, os.FileInfo) {
	// Use os.Stat to get file information
	info, err := os.Stat(filePath)

	// Check for errors
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("File does not exist."), nil
		} else {
			return err, nil
		}
	} else {
		return nil, info
	}
}
