package server

import "github.com/shawkyelshalawy/Daily_Brief/handlers"

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
}
