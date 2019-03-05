// +build !production

package config

const (
	// AdminSecret secret for check IsAdmin role
	AdminSecret = "i_am_admin"
	// DBConnectionString connection string to local postgresql
	DBConnectionString = "postgres://monita:monita@localhost:5432/monita?sslmode=disable"
	// Entrypoint first argument for ListenAndServe
	Entrypoint = ":8000"
	// JWTSecret for JWT Authorization
	JWTSecret = "jwt_secret"
)
