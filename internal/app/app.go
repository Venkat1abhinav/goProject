package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/venkat1abhinav/goProject/internal/api"
)

type Application struct {
	Logger  *log.Logger
	Workout *api.WorkOutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workOut := api.NewWorkoutHandler()

	app := &Application{
		Logger:  logger,
		Workout: workOut,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) ReturnFormmatedData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This will return a formated string\n")
}
