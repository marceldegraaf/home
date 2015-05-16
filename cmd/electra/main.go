package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/marceldegraaf/home/src/client"
	"github.com/marceldegraaf/home/src/devices/kaifa"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Infof("Electra starting")

	k := kaifa.Initialize()
	c := client.New()

	go k.Poll()

	for {
		select {
		case p := <-k.C:
			if err := c.Send(p); err != nil {
				log.Errorf("could not send data point: %s", err)
			}
		default:
		}
	}
}
