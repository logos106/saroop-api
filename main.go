package main

import (
	"log"
	"net/http"

  "github.com/rs/cors"

	"example.com/logos106/saroop-api/router"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", setupGlobalMiddleware(router)))
}
