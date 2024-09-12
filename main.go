package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/PullRequestInc/go-gpt3"
)

func main() {
	godotenv.load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API Key")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	response, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The capital of Suriname is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.Choices[0].Text)
}
