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
