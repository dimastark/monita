package user

import (
	"context"
	"log"

	"monita/storage/user"

	"github.com/go-chi/jwtauth"
)

// FromContext returns User struct from Claims
// returns `nil` if unauthorized or some error
func FromContext(ctx context.Context) *user.User {
	_, claims, err := jwtauth.FromContext(ctx)

	if err != nil {
		log.Println(err)

		return nil
	}

	if id, ok := claims["id"].(float64); ok {
		return user.GetByID(uint(id))
	}

	return nil
}
