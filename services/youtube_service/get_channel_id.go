package youtube_service

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// GetChannelIDsByChannelUrl fetches the YouTube page and extracts all channel IDs.
func GetChannelIDByChannelUrl(channelURL string) (string, error) {
	if !strings.HasSuffix(channelURL, "/videos") {
		channelURL = fmt.Sprintf("%s/videos", channelURL)
	}

	// Make an HTTP GET request to the channel URL.
	resp, err := http.Get(channelURL)
	if err != nil {
		return "", fmt.Errorf("failed to get URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	bodyStr := string(body)

	// logrus.Infof("bodyStr: %+v", bodyStr)

	var matches [][]string

	// Regular expression to find all channel IDs (e.g., "UCsXVk37bltHxD1rDPwtNM8Q").
	reID := regexp.MustCompile(`"browseId":"(UC[\w-]+)"`)
	matches = reID.FindAllStringSubmatch(bodyStr, -1)

	if len(matches) == 0 {
		// Regular expression to find all channel IDs (e.g., "UCsXVk37bltHxD1rDPwtNM8Q").
		reID := regexp.MustCompile(`"channelId":"(UC[\w-]+)"`)
		matches = reID.FindAllStringSubmatch(bodyStr, -1)
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("could not find any channel IDs")
	}

	// Extract just the channel ID part (the captured group).
	var ids []string
	for _, match := range matches {
		if len(match) > 1 {
			ids = append(ids, match[1])
		}
	}

	// logrus.Infof("MATCHES: %+v", ids)

	return ids[len(ids)-1], nil
}
