package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

/*
Product struct will contain information about VR experiences
*/
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

/*
Product slice contains catalog of VR experiences
*/
var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

/*
mySigningKey for the secret key
*/
var mySigningKey = []byte("secret")

/*
StatusHandler will be invoked when the user calls the /status route. it will return
a string with the message "API is up and running"
*/
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

/*
ProductHandler will be called when the user makes a GET Request to the /products endpoint or route.
This will return list of avialble for users to review
*/
var ProductHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// converting the slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

/*
AddFeedbackHandler will save feedback for a specific product.
*/
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte("Product Not Found"))
	}

})

/*
GetTokenHandler is issue JWT token
*/
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	token.Claims["admin"] = true
	token.Claims["name"] = "tomtom"
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	// Write the token to the browser window
	w.Write([]byte(tokenString))
})

/*
jwtMiddleware is to verify the jwt token
*/
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	// Instantiating the gorilla/mux router
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// serving static assets like images, css from the /static/{file} route
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Defining API routes
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", jwtMiddleware.Handler(ProductHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")
	r.Handle("/get-token", GetTokenHandler).Methods("GET")

	//Wrap LogginHandler from gorrila/handlers around our route. so that
	//logger is called first on each route request
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

// NotImplemented Handler, whenever an API is hit we will simply return
// the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
