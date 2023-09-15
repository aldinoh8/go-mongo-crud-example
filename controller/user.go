package controller

import (
	"example/dto"
	"example/helpers"
	"example/model"
	"example/repository"
	"example/services"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Repository repository.User
	Mailer     services.Mailer
}

func NewUserController(r repository.User, m services.Mailer) User {
	return User{Repository: r, Mailer: m}
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
	reqBody := dto.TransferReqBody{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to transfer", err.Error()))
	}

	user := ctx.Get("currentUser").(model.User)
	receiver, err := u.Repository.FindById(reqBody.ReceiverId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, helpers.GenerateErrorResponse("receiver not found", err.Error()))
	}

	err = u.Repository.Transfer(&user, &receiver, reqBody.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to transfer", err.Error()))
	}

	notificationMessage := fmt.Sprintf("successfully transfer %v to %s", reqBody.Amount, receiver.FullName)
	go u.Mailer.SendMail(user.Email, "transaction notification", notificationMessage)

	return ctx.JSON(http.StatusOK, "OK")
}

func (u User) DeleteAccount(ctx echo.Context) error {
	user := ctx.Get("currentUser").(model.User)
	err := u.Repository.DeleteAccount(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, helpers.GenerateErrorResponse("failed to delete account", err.Error()))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete account",
	})
}
