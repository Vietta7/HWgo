package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vietta7/HWgo/loms/internal/handler"
	"github.com/Vietta7/HWgo/loms/internal/repository"
	"github.com/Vietta7/HWgo/loms/internal/usecase"
)

func main() {
	repo := repository.NewMemoryRepo()
	uc := usecase.New(repo)
	h := handler.New(uc)

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Method string          `json:"method"`
			Params json.RawMessage `json:"params"`
			ID     any             `json:"id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := h.Handle(context.Background(), req.Method, req.Params)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	log.Println("loms listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}