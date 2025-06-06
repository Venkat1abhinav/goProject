package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/venkat1abhinav/goProject/internal/api"
	"github.com/venkat1abhinav/goProject/internal/migrations"
	"github.com/venkat1abhinav/goProject/internal/store"
)

type Application struct {
	Logger           *log.Logger
	ProductInventory *api.ProductInventoryHandler
	DB               *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, error := store.Open()
	if error != nil {
		return nil, error
	}
	error = store.MigrateFS(pgDB, migrations.FS, ".")

	if error != nil {
		panic(error)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//stores go here

	productStore := store.NewPostgresProductStore(pgDB)

	//handlers go here
	productInventory := api.NewProductHander(productStore)

	app := &Application{
		Logger:           logger,
		ProductInventory: productInventory,
		DB:               pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) ReturnFormmatedData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This will return a formated string\n")
}
