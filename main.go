package main

import (
	"database/sql"
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

import _ "github.com/lib/pq"

type apiConfig struct {
	Queries *database.Queries
}

func HandleApp(w http.ResponseWriter, r *http.Request) {
	component := ui.Index()
	component.Render(r.Context(), w)
}

func HandleBar(w http.ResponseWriter, r *http.Request) {
	component := ui.BarIndex()
	component.Render(r.Context(), w)
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to Database %s", err)
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
	fileSys := http.FileServer(http.Dir("./assets/"))
	strip := http.StripPrefix("/assets/", fileSys)

	serveMux.Handle("/assets/", strip)
	serveMux.Handle("/", http.HandlerFunc(HandleApp))
	serveMux.Handle("/bar", http.HandlerFunc(HandleBar))
	serveMux.Handle("GET /cockatils", http.HandlerFunc(cfg.getAllCocktails))
	serveMux.Handle("POST /api/addCocktail", http.HandlerFunc(cfg.addCocktails))
	serveMux.Handle("POST /api/addRecipie", http.HandlerFunc(cfg.addRecipieFunc))
	serveMux.Handle("GET /cock", http.HandlerFunc(cfg.getCocktailAPI))
	serveMux.Handle("POST /order", http.HandlerFunc(cfg.sendOrder))
	//TODO:
	// serveMux.Handle("POST /api/cancelOrder", http.HandlerFunc(cfg.cancelOrder))
	// serveMux.Handle("GET /api/myOrders", http.HandlerFunc(cfg.getAllCocktails))
	serveMux.Handle("GET /currentOrders", http.HandlerFunc(cfg.getCurrentOrders))
	serveMux.Handle("PUT /finishOrder", http.HandlerFunc(cfg.finishOrder))
	server.ListenAndServe()
}
