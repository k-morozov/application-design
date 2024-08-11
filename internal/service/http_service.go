package service

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/handlers"
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	ServiceHTTPOption func(s *ServiceHTTP)
	ServiceHTTP       struct {
		server *http.Server
		engine *chi.Mux
		store  storage.Storage
		config config.ServiceConfig
		log    zerolog.Logger
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

	guestHouseManager := guest_house.NewGuestHouseManager(srv.log)
	store, err := storage.NewStorage(guestHouseManager, srv.config, srv.log)
	if err != nil {
		srv.log.Err(err).Msg("failed create store")
		return nil, err
	}

	srv.store = store

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
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		s.log.Err(err)
		return err
	}

	log.Info().Msg("ListenAndServe is finishing")
	return nil
}

func (s *ServiceHTTP) Ping(rw http.ResponseWriter, req *http.Request) {
	handlers.Ping(rw, req, s.store)
}

func (s *ServiceHTTP) Orders(rw http.ResponseWriter, req *http.Request) {
	handlers.Orders(rw, req, s.store, s.config)
}
