package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sotanodroid/GO_API/pkg/api"
	"github.com/sotanodroid/GO_API/pkg/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {

	log.Println("Starting App...")

	models.InitDB(os.Getenv("DB_URL"))

	go api.RunServer()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Println("App stopping...")

	_, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunc()

	log.Println("App stopped")

}
