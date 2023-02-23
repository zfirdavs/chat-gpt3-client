package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		log.Fatalln("API key is missing")
	}

}
