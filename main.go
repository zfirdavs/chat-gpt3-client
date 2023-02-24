package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

const prompt = "I am a highly intelligent question answering bot. If you ask me a question that is rooted in truth, I will give you the answer. If you ask me a question that is nonsense, trickery, or has no clear answer, I will respond with \"Unknown\"\n"

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		log.Fatalln("API key is missing")
	}

	fmt.Println("To exit type q or quit")
	var question string
	ctx := context.Background()

	for {
		fmt.Println("Enter your question below:")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			question = strings.TrimSpace(scanner.Text())
		}

		switch question {
		case "":
			continue
		case "q", "quit":
			return
		}

		client := gpt3.NewClient(apiKey)

		resp, err := client.CompletionWithEngine(ctx, "davinci", gpt3.CompletionRequest{
			Prompt:           []string{prompt + "Q:" + question + "?\nA:"},
			MaxTokens:        gpt3.IntPtr(1000),
			Stop:             []string{"\n"},
			Echo:             false,
			Temperature:      gpt3.Float32Ptr(0),
			TopP:             gpt3.Float32Ptr(1),
			FrequencyPenalty: *gpt3.Float32Ptr(0),
			PresencePenalty:  *gpt3.Float32Ptr(0),
		})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(resp.Choices[0].Text)
	}
}
