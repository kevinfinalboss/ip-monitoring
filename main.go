package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kevinfinalboss/ip-monitoring/notifiers"
	"github.com/kevinfinalboss/ip-monitoring/routers"
	"github.com/kevinfinalboss/ip-monitoring/services"
	"github.com/sirupsen/logrus"
)

const discordWebhookUrl = "https://discord.com/api/webhooks/1103840164062707762/hmu05z5RrS4ya4QTHBKT7XxSaCfS1JxoACWZ750lzje0sZpejBY_6tu0AzK1pAshzJ4m"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		log.Println("Failed to log to file, using default stderr")
	}

	if os.Getenv("ENV") == "production" {
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	r := routers.NewRouter()

	go func() {
		urls, err := services.GetUrlsFromFile("urls.txt")
		if err != nil {
			log.Fatal(err)
		}

		for {
			for _, url := range urls {
				status, err := services.GetIPStatus(url)
				if err != nil {
					log.Println("Error getting status for", url, "-", err)
					continue
				}

				log.Println("Status for", url, "-", status)
				err = notifiers.PostToWebhook(discordWebhookUrl, status)
				if err != nil {
					log.Println("Error posting to discord webhook:", err)
				}
			}

			time.Sleep(1 * time.Hour)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", r))
}
