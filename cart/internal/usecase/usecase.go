package usecase

import "context"

type LomsClient interface {
	CheckStock(ctx context.Context, sku uint32, count uint16) (bool, error)
	CreateOrder(ctx context.Context, user int64, items []CartItem) (int64, error)
}

type ProductClient interface {
	GetProduct(ctx context.Context, sku uint32) (name string, price uint32, err error)
}

type CartRepository interface {
	AddItem(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteItem(ctx context.Context, user int64, sku uint32) error
	ListItems(ctx context.Context, user int64) ([]CartItem, error)
	Clear(ctx context.Context, user int64) error
}

type CartItem struct {
	Sku   uint32
	Count uint16
}

type CartListItem struct {
	Sku   uint32
	Count uint16
	Name  string
	Price uint32
}

type CartListResponse struct {
	Items      []CartListItem
	TotalPrice uint32
}

type CartUsecase interface {
	AddItem(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteItem(ctx context.Context, user int64, sku uint32) error
	List(ctx context.Context, user int64) (CartListResponse, error)
	Clear(ctx context.Context, user int64) error
	Checkout(ctx context.Context, user int64) (int64, error)
}

type cartUsecase struct {
	loms    LomsClient
	product ProductClient
	repo    CartRepository
}

func (c *cartUsecase) AddItem(ctx context.Context, user int64, sku uint32, count uint16) error {
    return nil
}

func (c *cartUsecase) DeleteItem(ctx context.Context, user int64, sku uint32) error {
    return nil
}

func (c *cartUsecase) List(ctx context.Context, user int64) (CartListResponse, error) {
    return CartListResponse{}, nil
}

func (c *cartUsecase) Clear(ctx context.Context, user int64) error {
    return nil
}

func (c *cartUsecase) Checkout(ctx context.Context, user int64) (int64, error) {
    return 0, nil
}


func New(loms LomsClient, product ProductClient, repo CartRepository) CartUsecase {
	return &cartUsecase{loms: loms, product: product, repo: repo}
}