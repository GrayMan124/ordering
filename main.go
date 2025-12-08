package main

import (
	"database/sql"
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

import _ "github.com/lib/pq"

type apiConfig struct {
	Queries *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
	dbQueries := database.New(db)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}
	cfg := apiConfig{Queries: dbQueries}
	serveMux := http.NewServeMux()
	server := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	fileSys := http.FileServer(http.Dir("."))
	strip := http.StripPrefix("/app", fileSys)
	serveMux.Handle("/app/", strip)
	serveMux.Handle("GET /api/cockatils", http.HandlerFunc(cfg.getAllCocktails))
	serveMux.Handle("POST /api/addCocktail", http.HandlerFunc(cfg.addCocktails))
	// TODO:
	// serveMux.Handle("POST /api/cockatil", http.HandlerFunc(cfg.getAllCocktails))
	serveMux.Handle("POST /api/order", http.HandlerFunc(cfg.sendOrder))
	serveMux.Handle("POST /api/cancelOrder", http.HandlerFunc(cfg.cancelOrder))
	// serveMux.Handle("GET /api/myOrders", http.HandlerFunc(cfg.getAllCocktails))
	serveMux.Handle("GET /api/currentOrders", http.HandlerFunc(cfg.getCurrentOrders))
	serveMux.Handle("PUT /api/FinishOrder", http.HandlerFunc(cfg.finishOrder))
	server.ListenAndServe()
}
