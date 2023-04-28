package chat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// Message represents a message sent to the OpenAI API.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request represents a request to the OpenAI API.
type Request struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// Response represents a response from the OpenAI API.
type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
}

// Choice represents a choice generated by the OpenAI API.
type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

// GenerateFromGPT generates a response from the OpenAI API.
func GenerateFromGPT(query string) (string, error) {
	req := Request{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: 300,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var resp Response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
