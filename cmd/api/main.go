package main

import (
	"log"

	"backend-service/config"
	"backend-service/docs"
	"backend-service/internal/api/controllers"
	"backend-service/internal/application/usecase"
	"backend-service/internal/infrastructure/database"
	"backend-service/internal/infrastructure/repositories"
	"backend-service/pkg/utilities/logger"
	"backend-service/pkg/utilities/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @contact.name				บริษัท โพ​ลาร์​ แบร์ ​มิ​ชชั่น ​จำกัด
// @contact.url				https://www.freshket.co/
// @contact.email				contact@freshket.co
//
// @title						Merchandise Pricing services
// @version					v1
// @description				Freshket Merchandise Pricing services API document
//
// @securityDefinitions.apikey	Authorization
// @in							header
// @name						Authorization
// @description				The token generated from Auth0
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.GetAppconfig()
	logger := logger.GetLogger()

	docs.SwaggerInfo.BasePath = "/" + conf.Basepath

	e := echo.New()

	e.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler(logger)

	dbClient := database.ConnectDB(conf.Client)

	repo := repositories.New(dbClient)
	usecase := usecase.New(repo)

	if conf.Env == "sit" || conf.Env == "uat" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	controllers.InitController(e, usecase)

	e.Logger.Fatal(e.Start(":" + conf.Port))
}
