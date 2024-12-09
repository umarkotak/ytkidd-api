package utils

import (
	"fmt"

	"github.com/umarkotak/ytkidd-api/config"
)

func GenFileUrl(guid string) string {
	return fmt.Sprintf("%s/ytkidd/api/file_bucket/%s", config.Get().AppHost, guid)
}
