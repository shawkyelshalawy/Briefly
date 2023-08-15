package server

import (
	"context"

	"github.com/shawkyelshalawy/Daily_Brief/handlers"
	"github.com/shawkyelshalawy/Daily_Brief/model"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux, s.database)

	handlers.FrontPage(s.mux)
	handlers.NewsletterSignup(s.mux, s.database, s.queue)
	handlers.NewsletterThanks(s.mux)
}

type signupperMock struct{}

func (s signupperMock) SignupForNewsletter(ctx context.Context, email model.Email) (string, error) {
	return "", nil
}
