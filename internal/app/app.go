package app

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/log"
	"applicationDesign/internal/service"
	"net/http"
)

func Run() {
	cfg := config.ParseConfig()
	lg := log.NewLogger(cfg.LogLevel)

	lg.Info().
		Any("cfg", cfg).
		Msg("get config")

	srv, err := service.NewServiceHTTP(lg, cfg,
		service.OptLogger(lg))

	if err != nil {
		lg.Error().Msg(err.Error())
		return
	}
	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		lg.Error().Err(err)
	}
}
