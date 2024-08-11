package service

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/handlers"
	"applicationDesign/internal/logic/hotel/manager"
	"applicationDesign/internal/provider"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	ServiceHTTPOption func(s *ServiceHTTP)
	ServiceHTTP       struct {
		server   *http.Server
		engine   *chi.Mux
		provider provider.Provider
		config   config.ServiceConfig
		log      zerolog.Logger
	}
)

func NewServiceHTTP(cfg config.ServiceConfig, opts ...ServiceHTTPOption) (*ServiceHTTP, error) {
	srv := &ServiceHTTP{
		engine: chi.NewRouter(),
		config: cfg,
	}

	for _, opt := range opts {
		opt(srv)
	}

	guestHouseManager := manager.NewGuestHouseManager(srv.log)
	serviceProvider, err := provider.NewProvider(guestHouseManager, srv.config, srv.log)
	if err != nil {
		srv.log.Err(err).Msg("failed create provider")
		return nil, err
	}

	srv.provider = serviceProvider

	return srv, nil
}

func (s *ServiceHTTP) ListenAndServe() error {
	s.engine.Get("/ping", s.Ping)
	s.engine.Post("/orders", s.Orders)

	log.Info().
		Str("port", s.config.Port).
		Msg("Starting server")

	log.Info().Msg("Using simple  HTTP")
	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.engine,
	}
	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		s.log.Err(err)
		return err
	}

	log.Info().Msg("ListenAndServe is finishing")
	return nil
}

func (s *ServiceHTTP) Ping(rw http.ResponseWriter, req *http.Request) {
	handlers.Ping(rw, req, s.provider)
}

func (s *ServiceHTTP) Orders(rw http.ResponseWriter, req *http.Request) {
	handlers.Orders(rw, req, s.provider, s.config)
}
