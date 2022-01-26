package main

import (
	"log"

	"github.com/joho/godotenv"

	"gitlab.i.bitonic.nl/lex/lnurl-server/server"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
	}

	server.Host()
}
