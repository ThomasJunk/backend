// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ThomasJunk/backend/configuration"
	"github.com/ThomasJunk/backend/db"
	"github.com/ThomasJunk/backend/hello"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Server struct {
	Router  *chi.Mux
	Session *scs.SessionManager
}

type App struct {
	Session       *scs.SessionManager
	Configuration *configuration.Configuration
	Logger        *zap.Logger
	Database      *db.Database
	HelloWorld    *hello.HelloWorld
}

func (app *App) initializeRessources() {
	app.HelloWorld = &hello.HelloWorld{
		Logger:   app.Logger,
		Database: app.Database,
	}

}

func Start(c *configuration.Configuration, l *zap.Logger, s *scs.SessionManager) {
	app := &App{
		Configuration: c,
		Logger:        l,
		Session:       s,
	}
	app.initializeRessources()
	srv := prepareServer(app, c)
	address := fmt.Sprintf("%s:%s", c.Address, c.Port)
	l.Info(fmt.Sprintf("Starting server on %s", address))
	l.Fatal(fmt.Sprintf("%s", srv.ListenAndServe()))
}

func prepareServer(app *App, c *configuration.Configuration) *http.Server {
	r := chi.NewRouter()
	srv := &Server{
		Router: r,
	}
	srv.AddMiddleWare(app)
	srv.AddRoutes(app)
	return &http.Server{
		Handler:      app.Session.LoadAndSave(srv.Router),
		Addr:         fmt.Sprintf("%s:%s", c.Address, c.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
