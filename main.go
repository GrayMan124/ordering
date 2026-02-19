package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/GrayMan124/ordering/server"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

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
	cfg := server.ApiConfig{Queries: dbQueries}
	serveMux := http.NewServeMux()
	server := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	fileSys := http.FileServer(http.Dir("./assets/"))
	strip := http.StripPrefix("/assets/", fileSys)

	serveMux.Handle("/assets/", strip)
	serveMux.Handle("/", http.HandlerFunc(cfg.Login))
	//TODO:
	// serveMux.Handle("POST /createUser", http.HandlerFunc(cfg.AddCocktail))
	serveMux.Handle("/bar", http.HandlerFunc(HandleBar))
	serveMux.Handle("GET /cockatils", http.HandlerFunc(cfg.GetAllCocktails))
	serveMux.Handle("GET /cock", http.HandlerFunc(cfg.GetCocktailData))
	serveMux.Handle("GET /leaderboard", http.HandlerFunc(cfg.LeaderBoardAPI))
	serveMux.Handle("POST /api/addCocktail", http.HandlerFunc(cfg.AddCocktail))
	serveMux.Handle("POST /AddCocktail", http.HandlerFunc(cfg.AddCocktailFromData))
	serveMux.Handle("GET /addRecipieForm", http.HandlerFunc(cfg.GetRecipieForm))
	serveMux.Handle("POST /order", http.HandlerFunc(cfg.SendOrder))
	serveMux.Handle("GET /currentOrders", http.HandlerFunc(cfg.GetCurrentOrders))
	serveMux.Handle("PUT /finishOrder", http.HandlerFunc(cfg.FinishOrder))
	//TODO:
	// serveMux.Handle("POST /api/cancelOrder", http.HandlerFunc(cfg.cancelOrder))
	// serveMux.Handle("GET /api/myOrders", http.HandlerFunc(cfg.getAllCocktails))
	server.ListenAndServe()
}
