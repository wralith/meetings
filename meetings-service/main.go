package main

import (
	"log"
	"net/http"
	"wra/meetings/pkg/api"
	"wra/meetings/pkg/repository"
	"wra/meetings/pkg/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func main() {
	a := App{}

	// Initializes Database Connections and API
	a.initialize()
	a.cors()

	// Initialize routes
	a.routes()

	log.Println("Server started at port 8080")
	err := http.ListenAndServe(":8080", a.Router)
	if err != nil {
		panic(err)
	}

}

func (a *App) cors() {
	a.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
}

func (a *App) routes() {
	API := InitAPI(a.DB)
	a.Router.Get("/meetings", API.FindAllMeetings())
	a.Router.Post("/meetings", API.CreateMeeting())
	a.Router.Get("/meetings/{id}", API.FindByID())
	a.Router.Put("/meetings/{id}", API.UpdateMeeting())
	a.Router.Delete("/meetings/{id}", API.DeleteMeeting())
}

func (a *App) initialize() {
	var err error
	const dbSource = "postgresql://postgres:123456@localhost:5432/meetings?sslmode=disable"

	a.DB, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")

	a.Router = chi.NewRouter()

}

func InitAPI(db *gorm.DB) api.MeetingsAPI {
	meetingsRepository := repository.New(db)
	meetingsService := service.New(meetingsRepository)
	meetingsAPI := api.New(*meetingsService)
	meetingsAPI.Migrate()
	return meetingsAPI
}
