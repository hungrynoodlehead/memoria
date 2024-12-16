package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/hungrynoodlehead/memoria/services/storage-service/handlers/photo_handler"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	_ "github.com/hungrynoodlehead/memoria/services/storage-service/docs"
)

type App struct {
	*chi.Mux
	Config  *utils.Config
	Logger  *utils.Logger
	DB      *utils.DB
	Storage *utils.Storage
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

	app.Mux = chi.NewMux()

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	app.Mount("/photo", photo_handler.NewPhotoHandler(app.Logger, app.DB, app.Storage, app.Config))

	err = http.ListenAndServe(":"+app.Config.GetListenPort(), app)
	if err != nil {
		return
	}
}
