package main

import (
	"github.com/gofiber/fiber/v2"
	"aziz/restshopcart/routes"
	"aziz/restshopcart/database"

)

func main(){

	app:= fiber.New()

	//INITIAL ROUTES
	routes.RoutInit(app)

	//initial database

	database.DatabaseInit()
	app.Listen(":3000")

}