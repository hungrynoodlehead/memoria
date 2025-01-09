package main

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/handlers/album_handler"
	"github.com/hungrynoodlehead/memoria/services/album-service/handlers/consumer_handler"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	_ "github.com/hungrynoodlehead/memoria/services/album-service/docs"
)

type Application struct {
	*echo.Echo
	DB     *utils.DB
	Config *utils.Config
}

// @Title Memoria Albums API
// @Version 0.1
// @License.name Apache 2.0
func main() {
	config, err := utils.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := utils.NewDB(config)
	if err != nil {
		panic(err)
	}

	producer, err := utils.NewMessageProducer(config)

	photoRepository := photo_repository.NewPhotoRepository(db, producer)
	albumRepository := album_repository.NewAlbumRepository(db, photoRepository)

	consumerHandler := consumer_handler.NewConsumerGroupHandler(config, db, photoRepository, albumRepository)
	err = consumer_handler.StartConsumer(config, consumerHandler)
	if err != nil {
		panic(err)
	}

	app := &Application{
		Echo: echo.New(),
		DB:   db,
	}

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	err = album_handler.BindAlbumHandler(app.Group("/album"), albumRepository, photoRepository)
	if err != nil {
		panic(err)
	}

	app.Logger.Fatal(app.Start(":8080"))
}
