package server

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type Server struct {
	name string
	addr string
	srv  *http.Server
}

func New(name, addr string, tlsConfig *tls.Config, handler http.Handler) *Server {
	s := &Server{name: name}
	s.srv = &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 2 * time.Second,
		TLSConfig:         tlsConfig,
	}
	return s
}

func (s *Server) Start(errCh chan error) {
	l, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		errCh <- err
		return
	}
	s.addr = l.Addr().String()

	slog.Info("listening", slog.String("name", s.name), slog.String("addr", l.Addr().String()))

	go func() {
		if s.srv.TLSConfig != nil {
			errCh <- s.srv.ServeTLS(l, "", "")
		} else {
			errCh <- s.srv.Serve(l)
		}
	}()
}

func (s *Server) Addr() string {
	return s.addr
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
