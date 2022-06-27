package objects

import (
	"net/http"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	m := request.Method
	if m == http.MethodPut {
		put(response, request)
		return
	}
	if m == http.MethodGet {
		get(response, request)
		return
	}
	response.WriteHeader(http.StatusMethodNotAllowed)
}
