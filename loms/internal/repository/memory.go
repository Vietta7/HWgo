package repository

import (
	"context"
	"sync"

	"github.com/Vietta7/HWgo/loms/internal/usecase"
)

type memoryRepo struct {
	mu     sync.Mutex
	nextID int64
}

func NewMemoryRepo() usecase.LomsRepository {
	return &memoryRepo{
		nextID: 1,
	}
}

func (m *memoryRepo) CreateOrder(ctx context.Context, user int64, items []usecase.OrderItem) (int64, error) {
	_ = ctx
	_ = user
	_ = items
	id := m.nextID
	m.nextID++
	return id, nil
}

func (m *memoryRepo) GetOrder(ctx context.Context, orderID int64) (usecase.Order, error) {
	_ = ctx
	return usecase.Order{
		ID:     orderID,
		User:   0,
		Status: usecase.StatusNew,
		Items:  nil,
	}, nil
}

func (m *memoryRepo) SetStatus(ctx context.Context, orderID int64, status usecase.OrderStatus) error {
	_ = ctx
	_ = orderID
	_ = status
	return nil
}

func (m *memoryRepo) ReserveStocks(ctx context.Context, items []usecase.OrderItem) (bool, error) {
	_ = ctx
	_ = items
	return true, nil
}

func (m *memoryRepo) ReleaseStocks(ctx context.Context, items []usecase.OrderItem) error {
	_ = ctx
	_ = items
	return nil
}

func (m *memoryRepo) CommitStocks(ctx context.Context, items []usecase.OrderItem) error {
	_ = ctx
	_ = items
	return nil
}

func (m *memoryRepo) StockInfo(ctx context.Context, sku uint32) (uint64, uint64, error) {
	_ = ctx
	_ = sku
	return 10, 0, nil
}