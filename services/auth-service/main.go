package main

import (
	"log"
	"net/http"

	_ "github.com/hungrynoodlehead/photos/services/auth-service/docs"
	authhandler "github.com/hungrynoodlehead/photos/services/auth-service/handlers/authHandler"
	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
	"github.com/hungrynoodlehead/photos/services/auth-service/middlewares"
	"github.com/hungrynoodlehead/photos/services/auth-service/utils"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Application struct {
	*chi.Mux
	DB     *gorm.DB
	Config *utils.Config
}

// @title						Swagger Example API
// @version					1.0
// @description				This is a sample server celler server.
// @termsOfService				http://swagger.io/terms/
//
// @Securitydefinitions.apikey	JWT Bearer
// @In							header
// @Name						Authorization
func main() {
	app := Application{Mux: chi.NewMux()}

	config, err := utils.NewConfig()

	if err != nil {
		log.Fatal("Cannot create app config: ", err)
		return
	} else {
		app.Config = config
	}

	db, err := utils.InitDatabase()
	if err != nil {
		log.Fatal("Cannot connect to Postgres: ", err)
		return
	} else {
		app.DB = db
	}

	_ = utils.InitLogger(config)

	jwtutils := jwtutils.NewJWTUtils(app.DB, app.Config)
	authMiddleware := middlewares.NewAuthMiddleware(app.DB, jwtutils)

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	app.Mount("/auth", authhandler.NewAuthHandler(app.DB, jwtutils, authMiddleware))

	http.ListenAndServe(":8080", app)
}
