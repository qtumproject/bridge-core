package service

import (
	"net"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/bridge/core/internal/config"
)

type service struct {
	log      *logan.Entry
	listener net.Listener
	cfg      config.Config
}

func (s *service) run() error {
	r := s.router()

	s.log.WithField("service", "api").Info("Service started")

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		listener: cfg.Listener(),
		cfg:      cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
