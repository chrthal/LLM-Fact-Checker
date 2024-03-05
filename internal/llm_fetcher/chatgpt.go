package llm_fetcher

import (
	"context"
	"fmt"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func GetResponse(query string) string {
	key := os.Getenv("OPENAI_KEY")

	client, errClient := chatgpt.NewClient(key)
	ctx := context.Background()
	if errClient != nil {
		fmt.Println("Error:", errClient)
	}
	res, err := client.SimpleSend(ctx, query)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return res.Choices[0].Message.Content
}
