package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/sotanodroid/GO_API/pkg/api"
	"github.com/sotanodroid/GO_API/pkg/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file")
	}

	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		log.Error(err)
	}

	log.SetLevel(log.Level(logLevel))
}

func main() {

	log.Info("Starting App...")

	models.InitDB(os.Getenv("DB_URL"))

	go api.RunServer()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Info("App stopping...")

	_, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunc()

	log.Info("App stopped")

}
