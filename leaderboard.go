package main

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"
)

func (cfg *apiConfig) LeaderBoardAPI(w http.ResponseWriter, r *http.Request) {
	rank, err := cfg.Queries.GetRanking(r.Context())
	if err != nil {
		log.Printf("Failed to get ranking from DB %v", err)
		w.WriteHeader(500)
		return
	}
	var names []string
	var scores []float32

	for _, position := range rank {
		names = append(names, position.OrderedBy)
		scores = append(scores, float32(position.Score))
	}

	w.WriteHeader(200)
	component := ui.LeaderBoard(names, scores)
	component.Render(r.Context(), w)
}
