package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/umarkotak/ytkidd-api/config"
)

func CreateFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory with appropriate permissions.
		// 0755 is a common permission setting (read/write/execute for owner, read/execute for others).
		err := os.MkdirAll(path, os.ModePerm) // Use MkdirAll to create parent directories if needed
		if err != nil {
			return fmt.Errorf("creating directory: %w", err)
		}
		fmt.Println("Folder created:", path)
	} else if err != nil {
		return fmt.Errorf("checking if directory exists: %w", err)
	} else {
		fmt.Println("Folder already exists:", path)
	}
	return nil
}

func GenFileUrl(guid string) string {
	return fmt.Sprintf("%s/ytkidd/api/file_bucket/%s", config.Get().AppHost, guid)
}

func GenRawFileUrl(bucketName, fileBucketPath string) string {
	return fmt.Sprintf("%s/%s/%s", config.Get().AppHost, bucketName, fileBucketPath)
}

func DeleteFileIfExists(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, which is fine.
			return nil
		}
		// Some other error occurred during deletion.
		return err
	}
	return nil
}

var errDangerousPath = errors.New("refusing to delete dangerous path")

// DeleteFolder deletes a directory and everything inside it.
// It refuses to delete "", ".", root, or the current working dir.
func DeleteFolder(path string) error {
	if path == "" || path == "." || path == "/" || path == `\` {
		return errDangerousPath
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Extra guard: don't allow deleting the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if abs == cwd || abs == string(filepath.Separator) {
		return errDangerousPath
	}

	// Optional: ensure it exists and is a directory (gives nicer errors)
	info, err := os.Lstat(abs)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("path is not a directory")
	}

	return os.RemoveAll(abs)
}
