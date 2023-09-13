package middleware

import (
	"example/helpers"
	"example/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	UserRepository repository.User
}

func NewAuthMiddleware(r repository.User) Auth {
	return Auth{UserRepository: r}
}

func (a Auth) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		token := ctx.Request().Header.Get("token")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		claims, err := helpers.VeriyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		user, err := a.UserRepository.FindById(claims["id"].(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token credentials")
		}

		ctx.Set("currentUser", user)
		return next(ctx)
	}
}
