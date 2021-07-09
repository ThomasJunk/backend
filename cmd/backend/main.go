// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ThomasJunk/backend/configuration"
	"github.com/ThomasJunk/backend/server"
	"github.com/alexedwards/scs/boltstore"
	"github.com/alexedwards/scs/v2"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func main() {
	c := configuration.New()
	l, _ := zap.NewProduction()
	boltdb := prepareSessionStore(l)
	s := prepareSessionManager(c, boltdb)
	defer l.Sync()
	server.Start(c, l, s)
}

func prepareSessionManager(config *configuration.Configuration, b *bbolt.DB) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Store = boltstore.NewWithCleanupInterval(b, 20*time.Second)
	sessionManager.Lifetime = 3 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	sessionManager.Cookie.Name = config.CookieName
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode
	sessionManager.Cookie.Secure = config.CookieSecure
	return sessionManager
}

func prepareSessionStore(l *zap.Logger) *bbolt.DB {
	boltdb, err := bbolt.Open("/tmp/bolt.db", 0600, nil)
	if err != nil {
		l.Error(fmt.Sprintf("%s", err))
	}
	return boltdb
}
