package handler

import (
	"encoding/json"
	"net/http"

	cache "product/internal/cache"
	store "product/internal/store"
)

type Handler struct {
	store *store.Store
	cache *cache.Cache
}

func NewHandler(s *store.Store, c *cache.Cache) *Handler {
	return &Handler{s, c}
}

// WRITE endpoint → Mutex usage
func (h *Handler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	h.store.IncrementView(id)

	w.Write([]byte("view incremented"))
}

// READ endpoint → RWMutex usage
func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// 1. Try cache
	if p, ok := h.cache.Get(id); ok {
		json.NewEncoder(w).Encode(p)
		return
	}

	// 2. Fetch from store (DB simulation)
	p, ok := h.store.GetProduct(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	// 3. Populate cache
	h.cache.Set(id, p)

	json.NewEncoder(w).Encode(p)
}