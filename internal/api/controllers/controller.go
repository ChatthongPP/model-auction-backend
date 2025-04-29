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
	group := e.Group("/api")
	group.POST("/hello", controller.Hello)

	group.GET("/categories", controller.GetCategories)
	group.POST("/register", controller.CreateUser)
}

func currentTime() *time.Time {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil
	}

	currentTime := time.Now().In(location)
	return &currentTime
}
