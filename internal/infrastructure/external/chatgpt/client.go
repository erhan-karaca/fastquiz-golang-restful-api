package chatgpt

import (
	"context"
	"fastquiz-api/pkg/config"
	"github.com/sashabaranov/go-openai"
	"log"
)

type Response struct {
	Content string
}

type Client struct {
	openaiClient *openai.Client
}

func NewChatGptClient() Client {
	client := openai.NewClient(config.AppConfig.ChatGptAPIKey)
	return Client{openaiClient: client}
}

func (c *Client) GenerateResponse(userPrompt string) (*Response, error) {
	req := newChatGptRequest(userPrompt)
	resp, err := c.openaiClient.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("API call error: %v", err)
	}

	content := resp.Choices[0].Message.Content

	return &Response{Content: content}, nil
}

func newChatGptRequest(prompt string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: config.AppConfig.ChatGptModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
}
