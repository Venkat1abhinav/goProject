package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/venkat1abhinav/goProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	r.Get("/products", app.ProductInventory.HandleCreateProductInventory)
	r.Get("/products/{id}", app.ProductInventory.HandleGetProductById)

	r.Post("/products", app.ProductInventory.HandleCreateProductInventory)
	r.Put("/products/{id}", app.ProductInventory.HandleUpdateProductInventory)

	r.Delete("/products/{id}", app.ProductInventory.HandleDeleteById)
	return r

}
