package client

import "context"


type LomsClient struct{}

func NewLomsClient(_ string) *LomsClient {
	return &LomsClient{}
}

func (c *LomsClient) CheckStock(ctx context.Context, sku uint32, count uint16) (bool, error) {
	_ = ctx
	_ = sku
	_ = count
	return true, nil
}

func (c *LomsClient) CreateOrder(ctx context.Context, user int64, items []struct {
	Sku   uint32
	Count uint16
}) (int64, error) {
	_ = ctx
	_ = user
	_ = items
	return 1, nil
}
