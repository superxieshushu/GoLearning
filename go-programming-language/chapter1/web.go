package chapter1

import (
	"log"
	"net/http"
)

func Web() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Lissajous(writer, request)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
