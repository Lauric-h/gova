package main

import (
	"gova/cmd"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
