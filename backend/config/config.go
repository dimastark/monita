// +build !production

package config

import (
	"os"
)

var (
	// AdminSecret secret for check IsAdmin role
	AdminSecret = "i_am_admin"
	// DBConnectionString connection string to local postgresql
	DBConnectionString = "postgres://monita:monita@localhost:5432/monita?sslmode=disable"
	// Entrypoint first argument for ListenAndServe
	Entrypoint = ":8000"
	// JWTSecret for JWT Authorization
	JWTSecret = "jwt_secret"
	// EmailFrom for sending emails
	EmailFrom = "robot-monita@yandex.ru"
	// DialerAddress for sending emails
	DialerAddress = "smtp.yandex.ru"
	// DialerPort for sending emails
	DialerPort = 465
	// DialerUser for sending emails
	DialerUser = "robot-monita"
	// DialerPassword for sending emails
	DialerPassword = os.Getenv("MONITA_DIALER_PASSWORD")
)
