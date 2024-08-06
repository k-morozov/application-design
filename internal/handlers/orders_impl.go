package handlers

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	"applicationDesign/internal/storage"
	"context"
)

func OrdersImpl(ctx context.Context, store storage.Storage, cfg config.ServiceConfig) error {
	lg := log.FromContext(ctx).With().Caller().Logger()

	lg.Debug().Msg("OrdersImpl started")

	if err := store.Orders(ctx); err != nil {
		lg.Err(err).
			Msg("failed got shortURL from store")
		return err
	}

	lg.Debug().Msg("OrdersImpl finished")

	return nil
}
