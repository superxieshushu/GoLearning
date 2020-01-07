package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func Server1() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var mu sync.Mutex
var count int

func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter2)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func Server3() {
	http.HandleFunc("/", handle3)
}

func handle3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "remoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q\n", k, v)
	}
}

func Server4() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Lissajous(writer)
	})
}

func P112() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		LissajousWithCycles(writer, request)
	})
}
