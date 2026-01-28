package client

import "context"

type ProductClient struct{}

func NewProductClient() *ProductClient {
	return &ProductClient{}
}

func (c *ProductClient) GetProduct(ctx context.Context, sku uint32) (string, uint32, error) {
	_ = ctx
	_ = sku
	return "stub product", 100, nil
}
