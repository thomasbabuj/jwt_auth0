package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instantiating the gorilla/mux router
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// serving static assets like images, css from the /static/{file} route
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Defining API routes
	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	http.ListenAndServe(":3000", r)
}

// NotImplemented Handler, whenever an API is hit we will simply return
// the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
