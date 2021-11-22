package router

import (

	"github.com/gofiber/fiber/v2"

   c "github.com/KeshikaGupta20/Postgresql_GO/controller"
)

func RegisterRoutes(app fiber.Router) {

	app.Post("/insertpro", c.AddEmployee)

	app.Delete("/deletepro/:id", c.DeleteEmployee)

	app.Get("/getpro", c.GetEmployee)

	app.Put("/updatepro/:id", c.GetEmployees)
}