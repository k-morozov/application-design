package handlers

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	RequestParser "applicationDesign/internal/parser"
	"applicationDesign/internal/provider"
	"context"
	"net/http"
)

func Orders(rw http.ResponseWriter, req *http.Request, serviceProvider provider.Provider, cfg config.ServiceConfig) {
	ctx, cancel := context.WithTimeout(req.Context(), cfg.HandleTimeout)
	defer cancel()
	lg := log.FromContext(ctx).With().Caller().Logger()

	lg.Debug().Msg("Orders handle started")

	order, err := RequestParser.ParseBodyOrderRequest(req, lg)
	if err != nil {
		lg.Error().Msg("Failed parse order from request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	lg.Debug().Any("Order parsed from request: %v", order)

	if err = serviceProvider.Orders(ctx, order); err != nil {
		lg.Err(err).
			Msg("failed got shortURL from store")
		http.Error(rw, "failed Orders", http.StatusInternalServerError)
		return
	}

	lg.Debug().Msg("finished")
	rw.WriteHeader(http.StatusCreated)
}
