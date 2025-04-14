package main

import (
	"fmt"
	"go-boilerplate/api"
	"go-boilerplate/internal/app"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()
	router := api.InitRouter(app)

	fmt.Printf("Server is running on http://%s\n", app.Address)
	log.Fatal(http.ListenAndServe(app.Address, router))
}