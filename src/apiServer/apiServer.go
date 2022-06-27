package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pineapple-man/flash-bird/src/apiServer/heartbeat"
	"github.com/pineapple-man/flash-bird/src/apiServer/locate"
	"github.com/pineapple-man/flash-bird/src/apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
