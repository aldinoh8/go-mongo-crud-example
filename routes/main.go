package routes

import (
	"example/config"
	"example/controller"
	"example/middleware"
	"example/repository"
	"example/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo) {
	db := config.InitDB()
	userRepository := repository.NewUserRepository(db)
	mailer := services.NewMailer("http://localhost:8001")

	userController := controller.NewUserController(userRepository, *mailer)
	auth := app.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	protected := app.Group("")
	authMiddleware := middleware.NewAuthMiddleware(userRepository)
	protected.Use(authMiddleware.Authenticate)
	{
		users := protected.Group("/users")
		users.GET("/detail", userController.Detail)
		users.POST("/topup", userController.Topup)
		users.DELETE("", userController.DeleteAccount)
	}
}
