package product

import (
	"sync"

	model "product/internal/model"
)

type Store struct {
	mu       sync.Mutex
	products map[string]*model.Product
}

func NewStore() *Store {
	return &Store{
		products: make(map[string]*model.Product),
	}
}

// WRITE → Mutex
func (s *Store) IncrementView(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p, ok := s.products[id]; ok {
		p.Views++
	}
}

func (s *Store) GetProduct(id string) (*model.Product, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.products[id]
	return p, ok
}

func (s *Store) Seed(products []*model.Product) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, p := range products {
		s.products[p.ID] = p
	}
}