package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	http.HandleFunc("/", index)
	port := getPort()
	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "nothing to see yet")
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	return port
}
