package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate the router outside of the main function
func newRouter() *mux.Router {
	// 1. Declare a new router
	// := is known as the short declaration operator.
	// It is used to declare and initialize the variables only inside the functions.
	// var is used to declare and initialize the variables inside and outside the functions.
	r := mux.NewRouter()
	r.HandleFunc("/hello", routeHandler).Methods("GET")

	// 2. Modifying the router to serve static files
	
	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above.
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}