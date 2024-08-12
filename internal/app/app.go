package app

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	"applicationDesign/internal/logic/rental/rental_manager"
	"applicationDesign/internal/service"
	"errors"
	"net/http"
)

func Run() {
	cfg := config.ParseConfig()
	lg := log.NewLogger(cfg.LogLevel)

	lg.Info().
		Any("cfg", cfg).
		Msg("get config")

	rentalManager := rental_manager.NewHotelManager(lg)
	srv, err := service.NewServiceHTTP(rentalManager, cfg,
		service.OptLogger(lg))

	if err != nil {
		lg.Error().Msg(err.Error())
		return
	}
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		lg.Error().Err(err)
	}
}
