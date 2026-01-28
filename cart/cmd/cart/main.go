package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
	
	"github.com/Vietta7/HWgo/cart/internal/client"
    "github.com/Vietta7/HWgo/cart/internal/handler"
    "github.com/Vietta7/HWgo/cart/internal/repository"
    "github.com/Vietta7/HWgo/cart/internal/usecase"
)

func main() {
	repo := repository.NewMemoryRepo()
	lomsClient := client.NewLomsClient("http://localhost:8082")
	productClient := client.NewProductClient()

	uc := usecase.New(lomsClient, productClient, repo)
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

		ctx := context.WithValue(r.Context(), handler.ContextKeyToken{}, r.Header.Get("Authorization"))
		resp := h.Handle(ctx, req.Method, req.Params)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	log.Println("cart listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}