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
	mailerService := services.NewMailer("http://localhost:8001")

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository, mailerService)
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
