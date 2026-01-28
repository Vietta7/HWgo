package handler

import (
	"context"
	"encoding/json"

	"github.com/Vietta7/HWgo/cart/internal/usecase"
)

type ContextKeyToken struct{}

type Handler struct {
	uc usecase.CartUsecase
}

func New(uc usecase.CartUsecase) *Handler {
	return &Handler{uc: uc}
}

type Response struct {
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	ID     any    `json:"id,omitempty"`
}

func (h *Handler) Handle(ctx context.Context, method string, params json.RawMessage) Response {
	switch method {
	case "cart.item.add":
		var req struct {
			User  int64  `json:"user"`
			Sku   uint32 `json:"sku"`
			Count uint16 `json:"count"`
		}
		_ = json.Unmarshal(params, &req)
		if err := h.uc.AddItem(ctx, req.User, req.Sku, req.Count); err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct{}{}}

	case "cart.item.delete":
		var req struct {
			User int64  `json:"user"`
			Sku  uint32 `json:"sku"`
		}
		_ = json.Unmarshal(params, &req)
		if err := h.uc.DeleteItem(ctx, req.User, req.Sku); err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct{}{}}

	case "cart.list":
		var req struct {
			User int64 `json:"user"`
		}
		_ = json.Unmarshal(params, &req)
		res, err := h.uc.List(ctx, req.User)
		if err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: res}

	case "cart.clear":
		var req struct {
			User int64 `json:"user"`
		}
		_ = json.Unmarshal(params, &req)
		if err := h.uc.Clear(ctx, req.User); err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct{}{}}

	case "cart.checkout":
		var req struct {
			User int64 `json:"user"`
		}
		_ = json.Unmarshal(params, &req)
		orderID, err := h.uc.Checkout(ctx, req.User)
		if err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct {
			OrderID int64 `json:"orderID"`
		}{OrderID: orderID}}

	default:
		return Response{Error: "unknown method"}
	}
}