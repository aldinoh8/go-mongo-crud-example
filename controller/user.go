package controller

import (
	"example/dto"
	"example/helper"
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
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	// validasi

	newUser, err := u.Repository.Create(&reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success register",
		"data":    newUser,
	})
}

func (u User) Login(ctx echo.Context) error {
	reqBody := dto.LoginBody{}

	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	// validasi

	// logic query find email pke mongo
	// 10baris

	// logic query find email postgre
	// 15baris

	user, err := u.Repository.FindByEmail(reqBody.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.ErrorResponse{Detail: "invalid email/password"})
	}

	isPasswordValid := helper.ComparePassword(reqBody.Password, user.Password)
	if !isPasswordValid {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.ErrorResponse{Detail: "invalid email/password"})
	}

	token := helper.GenerateToken(jwt.MapClaims{"id": user.ID})

	return ctx.JSON(http.StatusOK, echo.Map{"token": token})
}

func (u User) Detail(ctx echo.Context) error {
	user := ctx.Get("loggedinUser").(model.User)
	return ctx.JSON(http.StatusOK, user)
}

func (u User) Topup(ctx echo.Context) error {
	reqBody := dto.TopupBody{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	user := ctx.Get("loggedinUser").(model.User)
	user.DepositAmount = user.DepositAmount + reqBody.Amount

	err := u.Repository.UpdateAmount(user.ID, user.DepositAmount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	// send email
	u.Mailer.SendMail(
		user.Email,
		"Top up notification",
		fmt.Sprintf("Success topup: %v", user.DepositAmount),
	)

	// update amount menggunakan repository
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success topup",
		"data":    user,
	})
}

func (u User) DeleteAccount(ctx echo.Context) error {
	user := ctx.Get("loggedinUser").(model.User)
	err := u.Repository.DeleteAccount(user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.ErrorResponse{Detail: err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete account",
	})
}
