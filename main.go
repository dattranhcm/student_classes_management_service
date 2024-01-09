package main

import (
	"net/http"
	"student_classes_management_service/pkg/application/controller"
	dataaccess "student_classes_management_service/pkg/data-access"
	"student_classes_management_service/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	sqlDB := dataaccess.InitializeSequelDB("postgres://user:password@localhost:5432/student-service?sslmode=disable")
	userDA := dataaccess.NewUserDA(sqlDB)
	userService := service.NewUserService(userDA)
	userController := controller.NewUserController(userService)

	server := initializeHTTPServer()
	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	server.POST("/user", userController.CreateUser)
	server.GET("/users", userController.GetUsers)

	server.Logger.Fatal(server.Start("127.0.0.1:8080"))
}

func initializeHTTPServer() *echo.Echo {
	// Echo instance customization
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())  // Logger middleware
	e.Use(middleware.Recover()) // Panic recover middleware

	return e
}
