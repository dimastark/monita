// +build production

package config

import (
	"os"
)

const (
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
			panic()
		}
	}
}
