package main

import (
	"student_classes_management_service/pkg/application/constant"
	"student_classes_management_service/pkg/application/controller"
	dataaccess "student_classes_management_service/pkg/data-access"
	appMiddlewares "student_classes_management_service/pkg/middleware"
	"student_classes_management_service/pkg/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := initializeHTTPServer()
	// Error handler
	server.HTTPErrorHandler = appMiddlewares.ErrorHandler
	server.Validator = &appMiddlewares.CustomValidator{
		Validator: validator.New(),
	}

	sqlDB := dataaccess.InitializeSequelDB("postgres://user:password@localhost:5432/student-service?sslmode=disable")

	userRepo := dataaccess.NewUserRepo(sqlDB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	classRepo := dataaccess.NewClassRepo(sqlDB)
	classService := service.NewClassService(classRepo)
	classController := controller.NewClassController(classService)

	server.POST("/login", userController.Login)
	server.GET("/users", userController.GetUsers)

	authenticated := server.Group("/user", appMiddlewares.Authentication)
	authenticated.POST("/register", userController.CreateUser)
	authenticated.GET("/profile", userController.Profile)
	authenticated.GET("/course",userController.GetClasses)

	teacherRoute := server.Group("/class", appMiddlewares.Authentication,
		appMiddlewares.Authorization(constant.TeacherRole))

	teacherRoute.GET("/classes", classController.GetClasses)
	teacherRoute.POST("/class", classController.CreateClass)
	teacherRoute.POST("/classes/assign/:id", classController.AssignStudent)

	server.Logger.Fatal(server.Start("127.0.0.1:8080"))
}

func initializeHTTPServer() *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
