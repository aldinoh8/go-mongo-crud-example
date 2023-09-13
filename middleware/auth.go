package middleware

import "github.com/labstack/echo/v4"

type Auth struct {
}

func NewAuthMiddleware() Auth {
	return Auth{}
}

func (a Auth) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Auth Logic
		return next(ctx)
	}
}
