package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/venkat1abhinav/goProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/format", app.ReturnFormmatedData)
	r.Get("/workouts", app.Workout.CreateWorkOut)
	r.Get("/workouts/{id}", app.Workout.HandleWorkoutById)

	return r

}
