package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sotanodroid/GO_API/pkg/api"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	log.Println("Starting App...")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer conn.Close(context.Background())

	go api.RunServer()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Println("App stopping...")

	_, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunc()

	log.Println("App stopped")

}
