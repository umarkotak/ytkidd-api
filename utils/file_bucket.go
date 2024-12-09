package utils

import (
	"fmt"

	"github.com/umarkotak/ytkidd-api/config"
)

func GenFileUrl(guid string) string {
	return fmt.Sprintf("%s/ytkidd/api/file_bucket/%s", config.Get().AppHost, guid)
}

func GenRawFileUrl(fileBucketPath string) string {
	return fmt.Sprintf("%s/%s", config.Get().AppHost, fileBucketPath)
}
