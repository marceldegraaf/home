package server

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/marceldegraaf/home/src/storage"
)

func Start(port int) {
	log.Infof("Server listening on port :%d", port)

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/meter/measure", handleMeter)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing to see here.")
}

func handleMeter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "ERR: Method not allowed", http.StatusMethodNotAllowed)
	}

	p := storage.Point{}
	storage.Save(&p)

	fmt.Fprintf(w, "OK")
}
