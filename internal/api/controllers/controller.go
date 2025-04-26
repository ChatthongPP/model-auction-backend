package controllers

import (
	"time"

	"backend-service/internal/application/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	uc *usecase.Usecase
}

var validate *validator.Validate

func InitController(e *echo.Echo, usecase *usecase.Usecase) {
	controller := &Controller{uc: usecase}
	group := e.Group("/v1")
	group.POST("/hello", controller.Hello)
	group.POST("/login-line", controller.LoginLine)
	group.POST("/admins-management", controller.CreateAdmin)
	group.GET("/admins-management", controller.GetAdmins)
	group.PUT("/admins-management/:admins-id", controller.UpdateAdmin)
	group.DELETE("/admins-management/:admins-id", controller.DeleteAdmin)
	group.POST("/login", controller.AdminLogin)
}

func currentTime() *time.Time {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil
	}

	currentTime := time.Now().In(location)
	return &currentTime
}
