package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
}

func NewUserController() User {
	return User{}
}

func (u User) Register(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func (u User) Login(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func (u User) Detail(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func (u User) Transfer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
