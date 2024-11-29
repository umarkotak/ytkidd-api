package middlewares

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func decodeBasicAuth(auth string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return "", "", err
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid basic auth string: %s", auth)
	}

	return parts[0], parts[1], nil
}
