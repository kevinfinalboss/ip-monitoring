package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kevinfinalboss/ip-monitoring/routers"
	"github.com/sirupsen/logrus"
)

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

	log.Fatal(http.ListenAndServe(":8080", r))
}
