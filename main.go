package main

import (
	"log"
	"net/http"
	"time"
)

/*
9.

Channels are go's answer to this. Let's declare a channel that accepts strings
we'll give it a buffer of 1
*/
var stringChan = make(chan string, 1)

func main() {
	/*
		7.

		One of the main draws is concurrency, so let's try and demo that quickly
		the go keyword allows us to run any function in a different thread
	*/
	go doStuff()

	/*
		1.
		I'm going to show you how to build a quick and simple
		http server in go.

		The first thing we'll want to do is create a multiplexer

		It matches the URL of each incoming request against a list of registered
		patterns and calls the handler for the pattern that
		most closely matches the URL

		You've probably noticed that our new var is underlined red
		The go compiler will not compile if you have unused variables
	*/

	mux := http.NewServeMux()

	/*
		2.

		Next we'll give the multiplexer a route, and
		the name of the handler function associated with it

		mux is being used now, so it's no longer in red,
		but we haven't defined root, so we'll do that now
	*/
	mux.HandleFunc("/", root)

	/*
		4.

		Now we'll tell go to listen on an address and port
		We also need to pass in the multiplexer
	*/
	log.Println("starting server")
	http.ListenAndServe("localhost:8080", mux)
}

/*
3.

We have a root function declared, but the Handlefunc is still complaining, so what's up?

We can use intellij to show us what's expected.
Once we add the correct args, everything goes green again
*/
func root(w http.ResponseWriter, r *http.Request) {

	/*
		5.

		we can log anything we want about the request, so let's keep track of the route and http verb used
	*/
	log.Printf("received request a %s request from %s", r.Method, r.URL)

	/*
		11.

		now we can pass something to the channel, so let's send the remote address of the request
	*/
	stringChan <- r.RemoteAddr

	/*
		6.

		then let's just make this return something really simple for now
	*/
	w.Write([]byte("hello"))
}

/*
8.

If we create a function with an infinite loop that waits 5 and then logs something
if we run the server again, we can see that the first log is that the server has started
and that the new log happens every 5 seconds
we can still hit that endpoint, so that was easy, but how do we get these two threads to work together?
*/
func doStuff() {
	//for {
	//	time.Sleep(time.Second * 5)
	//	log.Println("doing stuff")
	//}

	/*
		10.

		Next let's rewrite this infinite loop to log strings that get passed in to the channel
		go gives us select statements for taking in data from channels. It looks just like a switch/case
		it picks the cases in random orders until the channel is closed
	*/
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
