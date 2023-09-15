package routes

import (
	"example/config"
	"example/controller"
	"example/middleware"
	"example/repository"

	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo) {
	db := config.InitDB()

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)
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
		users.POST("/transfer", userController.Transfer)
		users.DELETE("", userController.DeleteAccount)
	}
}
