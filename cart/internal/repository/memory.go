package repository

import (
	"context"
	"sync"

	"github.com/Vietta7/HWgo/cart/internal/usecase"
)

type memoryRepo struct {
	mu    sync.RWMutex
	carts map[int64]map[uint32]uint16
}

func NewMemoryRepo() usecase.CartRepository {
	return &memoryRepo{
		carts: make(map[int64]map[uint32]uint16),
	}
}

func (m *memoryRepo) AddItem(ctx context.Context, user int64, sku uint32, count uint16) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.carts[user]; !ok {
		m.carts[user] = make(map[uint32]uint16)
	}
	m.carts[user][sku] += count
	return nil
}

func (m *memoryRepo) DeleteItem(ctx context.Context, user int64, sku uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if cart, ok := m.carts[user]; ok {
		delete(cart, sku)
	}
	return nil
}

func (m *memoryRepo) ListItems(ctx context.Context, user int64) ([]usecase.CartItem, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var res []usecase.CartItem
	for sku, cnt := range m.carts[user] {
		res = append(res, usecase.CartItem{Sku: sku, Count: cnt})
	}
	return res, nil
}

func (m *memoryRepo) Clear(ctx context.Context, user int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.carts, user)
	return nil
}
