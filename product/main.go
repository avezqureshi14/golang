package main

import (
	"product/internal/cache"
	"product/internal/handler"
	model "product/internal/model"
	"product/internal/server"
	store "product/internal/store"
	"time"
)

func main() {
	store := store.NewStore()
	cache := cache.NewCache(1000, 5*time.Minute)

	store.Seed([]*model.Product{
		{ID: "1", Name: "iPhone"},
		{ID: "2", Name: "Laptop"},
	})

	h := handler.NewHandler(store, cache)

	server.Start(h)
}
