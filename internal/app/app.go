package app

import (
	"fmt"
	"github.com/TechGG1/Library/internal/handler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("loading env: %w", err)
	}

	//Start the server
	server := handler.NewServer()

	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown()

	<-done
	log.Printf("DONE!")

	return nil
}
