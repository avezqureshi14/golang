package handler

import (
	"encoding/json"
	"net/http"

	"golang.org/x/sync/singleflight"

	cache "product/internal/cache"
	store "product/internal/store"
)

type Handler struct {
	store *store.Store
	cache *cache.Cache
	group singleflight.Group
}

func NewHandler(s *store.Store, c *cache.Cache) *Handler {
	return &Handler{s, c, singleflight.Group{}}
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

	// 1. Cache hit
	if p, ok := h.cache.Get(id); ok {
		json.NewEncoder(w).Encode(p)
		return
	}

	// 2. singleflight (stampede protection)
	val, _, _ := h.group.Do(id, func() (interface{}, error) {

		// double check cache
		if p, ok := h.cache.Get(id); ok {
			return p, nil
		}

		// DB fetch (store)
		p, ok := h.store.GetProduct(id)
		if !ok {
			return nil, nil
		}

		h.cache.Set(id, p)
		return p, nil
	})

	if val == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(val)
}
