package handlers

import (
	"applicationDesign/internal/storage"
	"context"
)

func PingImpl(ctx context.Context, store storage.Storage) error {
	return store.Ping()
}
