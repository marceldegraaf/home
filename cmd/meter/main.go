package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/marceldegraaf/home/src/devices/kaifa"
)

func main() {
	log.Infof("Starting Kaifa smart meter poller")

	ch := kaifa.Initialize()

	for {
		select {
		case s := <-ch:
			log.Debugf("meter/main: received sample: %s", s)
		}
	}
}
