package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var msgGlob string

type echo struct {
	message chan string
}

func updateMsg(e *echo) {
	for {
		e.message <- msgGlob
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msgGlob = r.FormValue("w")
	log.Printf("echo from client")

	ec := &echo{
		message: make(chan string),
	}

	go updateMsg(ec)

	defer func() {
		close(ec.message)
		ec.message = nil
		log.Printf("client connection is closed")
	}()

	flusher := w.(http.Flusher)

	for {
		message := <-ec.message
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "word: %s\n", message)
		flusher.Flush()
	}
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msgGlob = r.FormValue("w")
}

func (e *echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	mux := http.NewServeMux()

	e := &echo{}

	mux.Handle("/events/", e)

	mux.HandleFunc("/echo", echoHandler)

	mux.HandleFunc("/say", sayHandler)

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", mux))
}
