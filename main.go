package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/kevinfinalboss/ip-monitoring/notifiers"
	"github.com/kevinfinalboss/ip-monitoring/routers"
	"github.com/kevinfinalboss/ip-monitoring/services"
	"github.com/sirupsen/logrus"
)

const (
	discordWebhookUrl = "https://discord.com/api/webhooks/1103840164062707762/hmu05z5RrS4ya4QTHBKT7XxSaCfS1JxoACWZ750lzje0sZpejBY_6tu0AzK1pAshzJ4m"
	port              = ":8080"
)

func main() {
	myFigure := figure.NewFigure("IP Monitoring", "", true)
	myFigure.Print()

	fmt.Println("API Name: IP Monitoring")
	fmt.Println("Version: 1.0.0")

	if os.Getenv("ENV") == "production" {
		logrus.SetLevel(logrus.WarnLevel)
		fmt.Println("Environment: Production")
	} else {
		logrus.SetLevel(logrus.DebugLevel)
		fmt.Println("Environment: Development")
	}

	fmt.Println("Listening on Port:", port)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Warnln("Failed to log to file, using default stderr")
	}

	r := routers.NewRouter()

	go func() {
		urls, err := services.GetUrlsFromFile("urls.txt")
		if err != nil {
			logrus.Fatalf("Error reading URLs from file: %v", err)
		}

		service := services.Service{}

		for {
			for _, url := range urls {
				status, err := service.GetIPStatus(url)
				if err != nil {
					logrus.Errorf("Error getting status for %s: %v", url, err)
					continue
				}

				logrus.Infof("Status for %s: %v", url, status)
				err = notifiers.PostToWebhook(discordWebhookUrl, status)
				if err != nil {
					logrus.Errorf("Error posting to discord webhook: %v", err)
				}
			}

			time.Sleep(1 * time.Hour)
		}
	}()

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	logrus.Infof("Server is ready to handle requests at %s", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	logrus.Infof("Server is shutting down due to %v signal", sig)

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Could not gracefully shutdown the server: %v\n", err)
	}
	logrus.Infof("Server stopped")
}
