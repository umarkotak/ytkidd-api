package utils

import (
	"fmt"
	"os"

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

func GenRawFileUrl(fileBucketPath string) string {
	return fmt.Sprintf("%s/%s", config.Get().AppHost, fileBucketPath)
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
