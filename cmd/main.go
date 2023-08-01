package cmd

import (
	"github.com/TechGG1/Library/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to connect with AnalystService: %s", err)
	}
}
