package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Response struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Index   int `json:"index"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

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

	response, err := http.DefaultClient.Do(request)
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

type Body struct {
	Message string `json:"message"`
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	engine := gin.New()
	engine.POST("/gogpt", func(context *gin.Context) {
		body := Body{}
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		gptmsg, err := GenerateFromGPT(body.Message)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		fmt.Printf("%s?\n%s\n", body.Message, gptmsg)
		context.JSON(http.StatusAccepted, &gptmsg)
	})
	engine.Run(":3000")
	fmt.Println("Running ...")
}
