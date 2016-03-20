package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/markstgodard/project-mayhem/durden"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	port := getPort()

	pm := durden.NewProjectMayhem()

	http.HandleFunc("/", pm.StatusHandler)

	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	return port
}
