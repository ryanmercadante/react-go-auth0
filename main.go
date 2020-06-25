package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instantiate gorilla/mux router
	r := mux.NewRouter()

	// On default page we will server static index page
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Setup server to serve static assets like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// API routes
	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	// App will run on port 8081
	http.ListenAndServe(":8081", r)
}

// NotImplemented is a handler for endpoints we haven't created yet
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
