package ollama_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
)

func SendChat(ctx context.Context, params ChatParams) (ChatResponse, error) {
	ollamaURL := fmt.Sprintf("%s/api/chat", config.Get().OllamaHost)

	params.Stream = false
	if params.Model == "" {
		params.Model = "hf.co/gmonsoon/gemma2-9b-cpt-sahabatai-v1-instruct-GGUF:Q5_0"
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return ChatResponse{}, err
	}

	req, err := http.NewRequest("POST", ollamaURL, bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return ChatResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return ChatResponse{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return ChatResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error call ollama api: %v", resp.StatusCode)
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"response_body": string(bodyBytes),
		}).Error(err)
		return ChatResponse{}, err
	}

	chatResponse := ChatResponse{}
	err = json.Unmarshal(bodyBytes, &chatResponse)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"response_body": string(bodyBytes),
		}).Error(err)
		return ChatResponse{}, err
	}

	return chatResponse, nil
}
