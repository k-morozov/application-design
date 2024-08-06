package handlers

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	"applicationDesign/internal/storage"
	"context"
	"net/http"
)

func Orders(rw http.ResponseWriter, req *http.Request, store storage.Storage, cfg config.ServiceConfig) {
	ctx, cancel := context.WithTimeout(req.Context(), cfg.HandleTimeout)
	defer cancel()
	lg := log.FromContext(ctx).With().Caller().Logger()

	lg.Debug().Msg("Orders handle started")

	if err := OrdersImpl(ctx, store, cfg); err != nil {
		http.Error(rw, "failed Orders", http.StatusInternalServerError)
		return
	}

	lg.Debug().Msg("finished")
	rw.WriteHeader(http.StatusOK)
}
