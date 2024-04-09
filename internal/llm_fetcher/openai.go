package llm_fetcher

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func gpt3_fetcher(question string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		log.Fatalln("ChatCompletion error: %v\n", err)
	}
	return resp.Choices[0].Message.Content
}

func gpt4_fetcher(question string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		log.Fatalln("ChatCompletion error: %v\n", err)
	}
	return resp.Choices[0].Message.Content
}
