package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Product is a new type that contains info about VR experiences
type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	{ID: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{ID: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{ID: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{ID: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{ID: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{ID: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

// NotImplemented is a handler for endpoints we haven't created yet
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

// StatusHandler will be invoked when the user calls the /status route
// It will return a string with the message "API is up and running!"
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running!"))
})

// ProductsHandler will be invoked when the user makes a GET request to the /products endpoint
// It will return a list of products available for users to review
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Convert slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// FeedbackHandler will add either positive or negative feedback to the product
// We would normally save this in a datbase
var FeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})

func main() {
	// Instantiate gorilla/mux router
	r := mux.NewRouter()

	// On default page we will server static index page
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Setup server to serve static assets like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// API routes
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", FeedbackHandler).Methods("POST")

	// App will run on port 8081
	http.ListenAndServe(":8081", r)
}
