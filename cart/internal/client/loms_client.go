package client

import (
	"context"

	"github.com/Vietta7/HWgo/cart/internal/usecase"
)


type LomsClient struct {
	addr string
}

func NewLomsClient(addr string) *LomsClient {
	return &LomsClient{addr: addr}
}

func (c *LomsClient) CheckStock(ctx context.Context, sku uint32, count uint16) (bool, error) {
	_ = ctx
	_ = sku
	_ = count
	return true, nil
}

func (c *LomsClient) CreateOrder(ctx context.Context, user int64, items []usecase.CartItem) (int64, error) {
	_ = ctx
	_ = user
	_ = items
	return 1, nil
}