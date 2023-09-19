package middleware

import (
	"example/dto"
	"example/helper"
	"example/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	UserRepository repository.User
}

func NewAuthMiddleware(u repository.User) Auth {
	return Auth{UserRepository: u}
}

func (a Auth) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Auth -> token
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.ErrorResponse{Detail: "invalid auth token"})
		}

		claims, err := helper.VeriyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.ErrorResponse{Detail: "invalid auth token"})
		}

		uid := claims["id"].(string)
		objectId, _ := primitive.ObjectIDFromHex(uid)
		user, err := a.UserRepository.FindById(objectId)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.ErrorResponse{Detail: "invalid auth token"})
		}

		ctx.Set("loggedinUser", user)
		return next(ctx)
	}
}
