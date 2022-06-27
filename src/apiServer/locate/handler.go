package locate

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method != http.MethodGet {

	}
	info := Locate(strings.Split(request.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	_, err := response.Write(b)
	if err != nil {
		log.Println(err)
		return
	}
}
