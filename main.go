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

	http.ListenAndServe(":3000", r)
}
