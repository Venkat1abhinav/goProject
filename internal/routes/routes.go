package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/venkat1abhinav/goProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/format", app.ReturnFormmatedData)
	r.Get("/products", app.ProductInventory.CreatProduct)
	r.Get("/products/{id}", app.ProductInventory.GetProductById)

	return r

}
