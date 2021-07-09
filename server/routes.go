// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package server

func (s *Server) AddRoutes(app *App) {
	s.Router.Get("/", app.HelloWorld.Welcome)
}
