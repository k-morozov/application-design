package handlers

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	RequestParser "applicationDesign/internal/parser"
	"applicationDesign/internal/provider"
	"context"
	"net/http"
)

func AddHotel(rw http.ResponseWriter, req *http.Request, serviceProvider provider.Provider, cfg config.ServiceConfig) {
	ctx, cancel := context.WithTimeout(req.Context(), cfg.HandleTimeout)
	defer cancel()
	lg := log.FromContext(ctx).With().Caller().Logger()

	lg.Debug().Msg("AddHotel handle started")

	newHotel, err := RequestParser.ParseBodyAddHotelRequest(req, lg)
	if err != nil {
		lg.Error().Msg("Failed parse order from request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	lg.Debug().Any("New hotel parsed from request: %v", newHotel)

	if err = serviceProvider.AddHotel(ctx, newHotel); err != nil {
		lg.Err(err).
			Msg("failed add hotel")
		http.Error(rw, "failed Orders", http.StatusInternalServerError)
		return
	}

	lg.Debug().Msg("finished")
	rw.WriteHeader(http.StatusCreated)
}
