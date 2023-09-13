package routes

import (
	"example/controller"
	"example/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo) {
	userController := controller.NewUserController()
	auth := app.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	protected := app.Group("")
	authMiddleware := middleware.NewAuthMiddleware()
	protected.Use(authMiddleware.Authenticate)
	{
		users := protected.Group("/users")
		users.GET("/detail", userController.Detail)
		users.POST("/transfer", userController.Transfer)
	}
}
