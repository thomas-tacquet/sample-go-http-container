package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("coucou"))
	})

	listeningPort := os.Getenv("LISTENING_PORT")

	if listeningPort == "" {
		listeningPort = ":8080"
	}

	if err := http.ListenAndServe(listeningPort, nil); err != nil {
		log.Panicf("error listen and serve on port %s", listeningPort)
	}
}
