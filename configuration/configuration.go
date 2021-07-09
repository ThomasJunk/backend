// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package configuration

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//Configuration holds server configuration.
type Configuration struct {
	Port           string
	Address        string
	CookieName     string
	CookieDomain   string
	CookieSecure   bool
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DB             string
	SSLMode        string
	Encryptioncost int
}

//New instantiates a new application configuration.
func New() *Configuration {
	err := godotenv.Load(".env.local")
	err = godotenv.Load()
	port := "3000"
	address := "127.0.0.1"
	cookiename := "session_id"
	cookiedomain := "example.com"
	cookiesecure := true
	dbhost := ""
	dbport := ""
	dbuser := ""
	dbpassword := ""
	db := ""
	sslmode := ""
	encryptioncost := 12
	if err == nil {
		port = os.Getenv("PORT")
		address = os.Getenv("ADDRESS")
		cookiename = os.Getenv("COOKIENAME")
		cookiedomain = os.Getenv("COOKIEDOMAIN")
		cookiesecure, _ = strconv.ParseBool(os.Getenv("COOKIESECURE"))
		dbhost = os.Getenv("DBHOST")
		dbport = os.Getenv("DBPORT")
		dbuser = os.Getenv("DBUSER")
		db = os.Getenv("DB")
		dbpassword = os.Getenv("DBPASSWORD")
		sslmode = os.Getenv("SSLMODE")
		encryptioncost, _ = strconv.Atoi(os.Getenv("ENCRYPTIONCOST"))
	}
	return &Configuration{
		Port:           port,
		Address:        address,
		CookieName:     cookiename,
		CookieDomain:   cookiedomain,
		CookieSecure:   cookiesecure,
		DBHost:         dbhost,
		DBPort:         dbport,
		DBUser:         dbuser,
		DBPassword:     dbpassword,
		DB:             db,
		SSLMode:        sslmode,
		Encryptioncost: encryptioncost,
	}
}
