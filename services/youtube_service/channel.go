package youtube_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ChannelDetails holds the information we want to return.
type ChannelDetails struct {
	Name         string
	ThumbnailURL string
	URL          string
}

// ---- Structs for parsing the API's JSON response ----

// ApiResponse matches the top-level structure of the YouTube API response.
type ApiResponse struct {
	Items []ChannelItem `json:"items"`
}

// ChannelItem represents a single channel resource in the response.
type ChannelItem struct {
	Snippet Snippet `json:"snippet"`
}

// Snippet contains the main details like title and thumbnails.
type Snippet struct {
	Title      string     `json:"title"`
	Thumbnails Thumbnails `json:"thumbnails"`
}

// Thumbnails contains URLs for various thumbnail sizes.
type Thumbnails struct {
	High Thumbnail `json:"high"`
}

// Thumbnail contains the URL for a specific thumbnail size.
type Thumbnail struct {
	URL string `json:"url"`
}

// GetYouTubeChannelDetails fetches details for a given channel ID using the YouTube Data API.
func GetYouTubeChannelDetails(apiKey, channelID string) (*ChannelDetails, error) {
	// 1. Construct the API request URL
	apiURL := "https://www.googleapis.com/youtube/v3/channels"
	params := url.Values{}
	params.Add("part", "snippet") // We need the 'snippet' to get title and thumbnails
	params.Add("id", channelID)
	params.Add("key", apiKey)

	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// 2. Make the HTTP GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %s: %s", resp.Status, string(body))
	}

	// 3. Decode the JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	// 4. Check if any channel was found
	if len(apiResponse.Items) == 0 {
		return nil, fmt.Errorf("no channel found for ID: %s", channelID)
	}

	// 5. Extract the data and populate our struct
	channelData := apiResponse.Items[0].Snippet

	details := &ChannelDetails{
		Name:         channelData.Title,
		ThumbnailURL: channelData.Thumbnails.High.URL,
		URL:          fmt.Sprintf("https://www.youtube.com/channel/%s", channelID),
	}

	return details, nil
}
