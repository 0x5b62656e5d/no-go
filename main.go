package main

import (
	"fmt"
	"log"
	"main/routes"
	"main/util"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var app *util.App

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found. Using existing process variables.")
	}

	app = &util.App{}

	app.LoadReasons()

	http.Handle("GET /health", util.RateLimitMiddleware(http.HandlerFunc(routes.HealthCheck)))
	http.Handle("GET /", util.RateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.No(w, r, app)
	})))

	port := os.Getenv("PORT")

	if port != "" {
		log.Printf("Listening on port %s...", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
		return
	}

	log.Printf("Listening on port 9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
