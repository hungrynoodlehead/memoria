package main

import (
	"encoding/json"
	"fmt"
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
	DB *utils.DB
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

	photoRepository := photo_repository.NewPhotoRepository(db)
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

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	albumGroup := app.Group("/album")
	albumGroup.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	err = album_handler.BindAlbumHandler(albumGroup, albumRepository, photoRepository)
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(app.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	app.Logger.Fatal(app.Start(":8080"))
}
