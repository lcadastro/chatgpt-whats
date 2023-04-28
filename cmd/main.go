package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	chat "github.com/lribeiros/whatsgpt"
)

type Body struct {
	Message string `json:"message"`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
}

func main() {
	engine := gin.New()
	engine.POST("/gogpt", func(context *gin.Context) {
		var body Body
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		gptmsg, err := chat.GenerateFromGPT(body.Message)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fmt.Printf("%s?\n%s\n", body.Message, gptmsg)
		context.JSON(http.StatusOK, gin.H{"message": gptmsg})
	})
	engine.Run(":3000")
}
