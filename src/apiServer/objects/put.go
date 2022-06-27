package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(response http.ResponseWriter, request *http.Request) {
	f, e := os.Open(os.Getenv("STORAGE_ROOT") +
		"/objects/" +
		strings.Split(request.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	_, err := io.Copy(f, request.Body)
	if err != nil {
		log.Println(err)
		return
	}

}
