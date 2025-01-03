package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/hungrynoodlehead/memoria/services/storage-service/handlers/photo_handler"
	"github.com/hungrynoodlehead/memoria/services/storage-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	_ "github.com/hungrynoodlehead/memoria/services/storage-service/docs"
)

type App struct {
	*chi.Mux
	Config          *utils.Config
	Logger          *utils.Logger
	DB              *utils.DB
	Storage         *utils.Storage
	Producer        *utils.BrokerProducer
	PhotoRepository *photo_repository.PhotoRepository
}

// @Title			Storage Microservice API
// @Version		0.1
// @license.name	Apache 2.0
func main() {
	var app App

	config, err := utils.NewConfig()
	if err != nil {
		panic(err)
	} else {
		app.Config = config
	}

	log, err := utils.NewLogger(app.Config)
	if err != nil {
		panic(err)
	} else {
		app.Logger = log
	}

	db, err := utils.NewDB(app.Config, app.Logger)
	if err != nil {
		panic(err)
	} else {
		app.DB = db
	}

	storage, err := utils.NewStorage(app.Config, app.Logger)
	if err != nil {
		panic(err)
	} else {
		app.Storage = storage
	}

	producer, err := utils.NewBrokerProducer(app.Logger, app.Config)
	if err != nil {
		panic(err)
	} else {
		app.Producer = producer
	}

	photoRepository := photo_repository.NewPhotoRepository(app.DB, app.Storage, app.Logger, app.Producer)
	app.PhotoRepository = photoRepository

	app.Mux = chi.NewMux()

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	app.Mount("/photos", photo_handler.NewPhotoHandler(app.Logger, app.DB, app.Storage, app.Config, app.PhotoRepository))

	err = http.ListenAndServe(":"+app.Config.GetListenPort(), app)
	if err != nil {
		return
	}
}
