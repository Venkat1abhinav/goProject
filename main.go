package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/venkat1abhinav/goProject/internal/app"
	"github.com/venkat1abhinav/goProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	app, error := app.NewApplication()

	if error != nil {
		panic(error)
	}
	defer app.DB.Close()
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("We are running on Port %d\n", port)

	err := server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
