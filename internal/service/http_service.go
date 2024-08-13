package service

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/handlers"
	"applicationDesign/internal/logic/rental/rental_manager"
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

func NewServiceHTTP(rentalManager rental_manager.BaseRentalManager, cfg config.ServiceConfig, opts ...ServiceHTTPOption) (*ServiceHTTP, error) {
	srv := &ServiceHTTP{
		engine: chi.NewRouter(),
		config: cfg,
	}

	for _, opt := range opts {
		opt(srv)
	}

	serviceProvider, err := provider.NewProvider(rentalManager, srv.config, srv.log)
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
	s.engine.Post("/add_hotel", s.AddHotel)

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

func (s *ServiceHTTP) AddHotel(rw http.ResponseWriter, req *http.Request) {
	handlers.AddHotel(rw, req, s.provider, s.config)
}
