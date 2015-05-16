package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/marceldegraaf/home/src/server"
)

const PORT = 8080

func main() {
	log.Info("Zeus starting")
	server.Start(PORT)
}
