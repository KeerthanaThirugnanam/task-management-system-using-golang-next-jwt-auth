package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func GetTaskSuggestions(c *gin.Context) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.TODO(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "Suggest a task for a user."},
			},
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI request failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suggestion": resp.Choices[0].Message.Content})
}
