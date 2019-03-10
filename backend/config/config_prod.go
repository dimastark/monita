// +build production

package config

import (
	"log"
	"os"
)

var (
	AdminSecret        = os.Getenv("MONITA_ADMIN_SECRET")
	DBConnectionString = os.Getenv("MONITA_DB_CONNECTION_STRING")
	Entrypoint         = os.Getenv("MONITA_ENTRYPOINT")
	JWTSecret          = os.Getenv("MONITA_JWT_SECRET")
	EmailFrom          = os.Getenv("MONITA_EMAIL_FROM")
	DialerAddress      = os.Getenv("MONITA_DIALER_ADDRESS")
	DialerPort         = os.Getenv("MONITA_DIALER_PORT")
	DialerUser         = os.Getenv("MONITA_DIALER_USER")
	DialerPassword     = os.Getenv("MONITA_DIALER_PASSWORD")
)

func init() {
	vars := []string{
		"MONITA_ADMIN_SECRET",
		"MONITA_DB_CONNECTION_STRING",
		"MONITA_ENTRYPOINT",
		"MONITA_JWT_SECRET",
		"MONITA_EMAIL_FROM",
		"MONITA_DIALER_ADDRESS",
		"MONITA_DIALER_PORT",
		"MONITA_DIALER_USER",
		"MONITA_DIALER_PASSWORD",
	}

	for _, v := range vars {
		if os.Getenv(v) == "" {
			log.Fatalln(v, "environment variable is not set")
		}
	}
}
