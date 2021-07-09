// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package server

import (
	"net/http"

	chilogger "github.com/766b/chi-logger"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

func (s *Server) AddMiddleWare(app *App) {
	s.Router.Use(chilogger.NewZapMiddleware("router", app.Logger))
	s.Router.Use(middleware.Logger)
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})
	s.Router.Use(secureMiddleware.Handler)
}

//Authenticated secures routes.
func (s *Server) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := s.Session.GetString(r.Context(), "userid")
		if msg == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
