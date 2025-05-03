package controllers

import (
	"backend-service/internal/application/usecase"
	"backend-service/pkg/utilities/middlewares"

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

	// Public Routes
	group.POST("/hello", controller.Hello)
	group.POST("/register", controller.CreateUser)
	group.POST("/login", controller.Login)
	group.GET("/categories", controller.GetCategories)
	group.GET("/products", controller.GetProducts)

	authGroup := group.Group("")
	authGroup.Use(middlewares.AuthMiddleware)
	authGroup.GET("/me", controller.GetProfile)
	authGroup.GET("/users/:id", controller.GetUserByID)
	authGroup.POST("/products", controller.CreateProduct)
	authGroup.GET("/products/:id", controller.GetProductByID)
	authGroup.POST("/bids", controller.CreateBid)
	authGroup.GET("/bids", controller.GetBids)
	authGroup.GET("/wallet-logs", controller.GetWalletLogs)
}

// func currentTime() *time.Time {
// 	location, err := time.LoadLocation("Asia/Bangkok")
// 	if err != nil {
// 		return nil
// 	}

// 	currentTime := time.Now().In(location)
// 	return &currentTime
// }
