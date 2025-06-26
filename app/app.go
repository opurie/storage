package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"main/app/handler"
	model "main/app/models"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(c *config.Config) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		c.DB.Username,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
		c.DB.Charset)
	db, err := gorm.Open(c.DB.Dialect, dbUri)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

	log.Println("App initialized successfully")
}

func (a *App) setRouters() {
	// Items
	a.Get("/items", a.handleRequest(handler.GetItems))
	a.Get("/items/{id}", a.handleRequest(handler.GetItem))
	a.Put("/items/{name}/{owner_id}/{category_id}/{location_id}/{tags}", a.handleRequest(handler.AddItem))
	a.Delete("/items/{id}", a.handleRequest(handler.DeleteItem))
	// Users
	a.Get("/users", a.handleRequest(handler.GetUsers))
	a.Get("/users/{id}", a.handleRequest(handler.GetUser))
	a.Put("/users/{username}/{email}", a.handleRequest(handler.AddUser))
	a.Post("/users/{id}/{username}/{email}", a.handleRequest(handler.UpdateUser))
	a.Delete("/users/{id}", a.handleRequest(handler.DeleteUser))
	// Locations
	a.Get("/locations", a.handleRequest(handler.GetLocations))
	a.Get("/locations/{id}", a.handleRequest(handler.GetLocation))
	a.Put("/locations/{name}", a.handleRequest(handler.AddLocation))
	a.Delete("/locations/{id}", a.handleRequest(handler.DeleteLocation))
	// Categories
	a.Get("/categories", a.handleRequest(handler.GetCategories))
	a.Get("/categories/{id}", a.handleRequest(handler.GetCategory))
	a.Put("/categories/{name}", a.handleRequest(handler.AddCategory))
	a.Delete("/categories/{id}", a.handleRequest(handler.DeleteCategory))
	// Users
	// Items
	// Locations
	// Categories
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
