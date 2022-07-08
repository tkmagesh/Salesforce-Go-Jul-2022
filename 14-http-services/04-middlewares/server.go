package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte("Foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("Bar"))
}

/*
func logger(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		handler(w, r)
		log.Println("Request processed")
	}
}

func profile(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			elapsed := end.Sub(start) / time.Millisecond
			fmt.Printf("%q took %dms\n", r.URL.Path, elapsed)
		}()
		handler(w, r)
	}
}
*/

func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		handler(w, r)
		log.Println("Request processed")
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			elapsed := end.Sub(start) / time.Millisecond
			fmt.Printf("%q took %dms\n", r.URL.Path, elapsed)
		}()
		handler(w, r)
	}
}

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func main() {
	/*
		http.HandleFunc("/foo", foo)
		http.HandleFunc("/bar", bar)
	*/

	/*
		http.HandleFunc("/foo", logger(foo))
		http.HandleFunc("/bar", logger(bar))
	*/

	/*
		http.HandleFunc("/foo", profile(foo))
		http.HandleFunc("/bar", profile(bar))
	*/

	/*
		http.HandleFunc("/foo", profile(logger(foo)))
		http.HandleFunc("/bar", profile(logger(bar)))
	*/

	http.HandleFunc("/foo", chain(foo, logger, profile))
	http.HandleFunc("/bar", chain(bar, logger, profile))
	http.ListenAndServe(":8080", nil)
}
