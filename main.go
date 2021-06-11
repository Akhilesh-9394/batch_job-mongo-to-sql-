package main

import (
	"batch_job/services"

	_ "github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	gocron.Every(30).Seconds().Do(services.Getdata, int64(10))
	<-gocron.Start()
}
