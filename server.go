package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return http.Server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return http.Server.Shutdown(ctx)

}
