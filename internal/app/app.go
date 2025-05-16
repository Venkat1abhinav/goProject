package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/venkat1abhinav/goProject/internal/api"
	"github.com/venkat1abhinav/goProject/internal/store"
)

type Application struct {
	Logger  *log.Logger
	Workout *api.WorkOutHandler
	DB      *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workOut := api.NewWorkoutHandler()
	pgDB, error := store.Open()

	if error != nil {
		return nil, error
	}

	app := &Application{
		Logger:  logger,
		Workout: workOut,
		DB:      pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) ReturnFormmatedData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This will return a formated string\n")
}
