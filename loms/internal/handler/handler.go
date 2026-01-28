package handler

import (
	"context"
	"encoding/json"

	"github.com/Vietta7/HWgo/loms/internal/usecase"
)

type Handler struct {
	uc usecase.LomsUsecase
}

func New(uc usecase.LomsUsecase) *Handler {
	return &Handler{uc: uc}
}

type Response struct {
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	ID     any    `json:"id,omitempty"`
}

func (h *Handler) Handle(ctx context.Context, method string, params json.RawMessage) Response {
	switch method {
	case "order/create":
		var req struct {
			User  int64 `json:"user"`
			Items []struct {
				Sku   uint32 `json:"sku"`
				Count uint16 `json:"count"`
			} `json:"items"`
		}
		_ = json.Unmarshal(params, &req)

		var items []usecase.OrderItem
		for _, it := range req.Items {
			items = append(items, usecase.OrderItem{Sku: it.Sku, Count: it.Count})
		}

		id, err := h.uc.OrderCreate(ctx, req.User, items)
		if err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct {
			OrderID int64 `json:"orderID"`
		}{OrderID: id}}

	case "order/info":
		var req struct {
			OrderID int64 `json:"orderID"`
		}
		_ = json.Unmarshal(params, &req)

		o, err := h.uc.OrderInfo(ctx, req.OrderID)
		if err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: o}

	case "order/pay":
		var req struct {
			OrderID int64 `json:"orderID"`
		}
		_ = json.Unmarshal(params, &req)

		if err := h.uc.OrderPay(ctx, req.OrderID); err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct{}{}}

	case "order/cancel":
		var req struct {
			OrderID int64 `json:"orderID"`
		}
		_ = json.Unmarshal(params, &req)

		if err := h.uc.OrderCancel(ctx, req.OrderID); err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct{}{}}

	case "stock/info":
		var req struct {
			Sku uint32 `json:"sku"`
		}
		_ = json.Unmarshal(params, &req)

		count, err := h.uc.StockInfo(ctx, req.Sku)
		if err != nil {
			return Response{Error: err.Error()}
		}
		return Response{Result: struct {
			Count uint64 `json:"count"`
		}{Count: count}}

	default:
		return Response{Error: "unknown method"}
	}
}