// +build production

package config

import (
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
		AdminSecret,
		DBConnectionString,
		Entrypoint,
		JWTSecret,
	}

	for _, v := range vars {
		if v == "" {
			panic("Not all environment variables set (MONITA_ADMIN_SECRET, MONITA_DB_CONNECTION_STRING, MONITA_ENTRYPOINT, MONITA_JWT_SECRET)")
		}
	}
}
