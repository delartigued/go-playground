package main

import (
	"log"
	"net/http"
	"time"
)

var stringChan = make(chan string, 1)

func main() {
	go doStuff()
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	log.Println("starting server")
	http.ListenAndServe("localhost:8081", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request from %s", r.URL)
	stringChan <- r.RemoteAddr
	w.Write([]byte("hello"))
}

func doStuff() {
	for {
		select {
		case receivedString := <-stringChan:
			log.Println("Got:", receivedString)
		default:
			time.Sleep(time.Second * 5)
			log.Println("doing stuff")
		}
	}
}
