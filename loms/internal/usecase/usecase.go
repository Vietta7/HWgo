package usecase

import "context"

type OrderStatus string

const (
	StatusNew             OrderStatus = "new"
	StatusAwaitingPayment OrderStatus = "awaiting payment"
	StatusFailed          OrderStatus = "failed"
	StatusPayed           OrderStatus = "payed"
	StatusCancelled       OrderStatus = "cancelled"
)

type OrderItem struct {
	Sku   uint32
	Count uint16
}

type Order struct {
	ID     int64
	User   int64
	Status OrderStatus
	Items  []OrderItem
}

type LomsRepository interface {
	CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error)
	GetOrder(ctx context.Context, orderID int64) (Order, error)
	SetStatus(ctx context.Context, orderID int64, status OrderStatus) error
	ReserveStocks(ctx context.Context, items []OrderItem) (bool, error)
	ReleaseStocks(ctx context.Context, items []OrderItem) error
	CommitStocks(ctx context.Context, items []OrderItem) error
	StockInfo(ctx context.Context, sku uint32) (total uint64, reserved uint64, err error)
}

type LomsUsecase interface {
	OrderCreate(ctx context.Context, user int64, items []OrderItem) (int64, error)
	OrderInfo(ctx context.Context, orderID int64) (Order, error)
	OrderPay(ctx context.Context, orderID int64) error
	OrderCancel(ctx context.Context, orderID int64) error
	StockInfo(ctx context.Context, sku uint32) (uint64, error)
}

func (u *lomsUsecase) OrderCreate(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	_ = ctx
	_ = user
	_ = items
	return u.repo.CreateOrder(ctx, user, items)
}

func (u *lomsUsecase) OrderInfo(ctx context.Context, orderID int64) (Order, error) {
	return u.repo.GetOrder(ctx, orderID)
}

func (u *lomsUsecase) OrderPay(ctx context.Context, orderID int64) error {
	_ = orderID
	return nil
}

func (u *lomsUsecase) OrderCancel(ctx context.Context, orderID int64) error {
	_ = orderID
	return nil
}

func (u *lomsUsecase) StockInfo(ctx context.Context, sku uint32) (uint64, error) {
	total, reserved, err := u.repo.StockInfo(ctx, sku)
	if err != nil {
		return 0, err
	}
	if total < reserved {
		return 0, nil
	}
	return total - reserved, nil
}

type lomsUsecase struct {
	repo LomsRepository
}

func New(repo LomsRepository) LomsUsecase {
	return &lomsUsecase{repo: repo}
}