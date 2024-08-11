package app

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
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

	srv, err := service.NewServiceHTTP(cfg,
		service.OptLogger(lg))

	if err != nil {
		lg.Error().Msg(err.Error())
		return
	}
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		lg.Error().Err(err)
	}
}
