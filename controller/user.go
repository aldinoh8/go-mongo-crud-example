package controller

import (
	"example/dto"
	"example/helpers"
	"example/model"
	"example/repository"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Repository repository.User
}

func NewUserController(r repository.User) User {
	return User{Repository: r}
}

func (u User) Register(ctx echo.Context) error {
	reqBody := dto.RegisterBody{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to create user", err.Error()))
	}

	newUser, err := u.Repository.Register(reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to create user", err.Error()))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success register",
		"newUser": newUser,
	})
}

func (u User) Login(ctx echo.Context) error {
	reqBody := dto.LoginBody{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to login user", err.Error()))
	}

	user, err := u.Repository.FindByEmail(reqBody.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, helpers.GenerateErrorResponse("failed to login user", "invalid email/password"))
	}

	validPassword := helpers.ComparePassword(reqBody.Password, user.Password)
	if !validPassword {
		return echo.NewHTTPError(http.StatusUnauthorized, helpers.GenerateErrorResponse("failed to login user", "invalid email/password"))
	}

	token := helpers.GenerateToken(jwt.MapClaims{"id": user.ID})

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (u User) Detail(ctx echo.Context) error {
	user := ctx.Get("currentUser").(model.User)
	return ctx.JSON(http.StatusOK, user)
}

func (u User) Transfer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
