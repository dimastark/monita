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
)

func init() {
	vars := []string{
		"MONITA_ADMIN_SECRET",
		"MONITA_DB_CONNECTION_STRING",
		"MONITA_ENTRYPOINT",
		"MONITA_JWT_SECRET",
	}

	for _, v := range vars {
		if os.Getenv(v) == "" {
			log.Fatalln(v, "environment variable is not set")
		}
	}
}
