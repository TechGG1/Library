package app

import (
	"fmt"
	"github.com/TechGG1/Library/internal/handler"
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/repository"
	"github.com/TechGG1/Library/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("loading env: %w", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		return fmt.Errorf("create db: %w", err)
	}
	defer db.Close()

	defaultLogLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return fmt.Errorf("log lvl: %w", err)
	}
	loggging := logging.NewLogger(defaultLogLevel)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, loggging)
	handlers := handler.NewHandler(services, *loggging)

	//Start the server
	server := handler.NewServer(os.Getenv("API_SERVER_HOST"), handlers.InitRoutes())

	done := make(chan bool)
	go func() {
		err := server.Run()
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
