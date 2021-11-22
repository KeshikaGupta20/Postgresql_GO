package controller

import (
	"fmt"

	"github.com/KeshikaGupta20/Postgresql_GO/database"
	"github.com/KeshikaGupta20/Postgresql_GO/models"

	"github.com/gofiber/fiber/v2"
)

func Print() {
	fmt.Println("Error ")
}

func AddEmployee(c *fiber.Ctx) error {
	fmt.Println("error")
	db := database.DB

	emp := new(models.Employee)

	c.BodyParser(emp)

	db.Create(&emp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Employee added sucessfully",
	})

}

func DeleteEmployee(c *fiber.Ctx) error {

	ID := c.Params("empid")

	db := database.DB

	var employee []models.Employee = make([]models.Employee, 0)

	db.Find(c.Context(), &employee, ID)

	db.Delete(&employee)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Product deleted sucessfully",
	})
}

//to reterive all employee details
func GetEmployees(c *fiber.Ctx) error {

	db := database.DB

	var employee []models.Employee = make([]models.Employee, 0)

	db.Find(c.Context(), &employee)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"data": fiber.Map{
			"Emp": "employee",
		},
	})

}

//to reterive single employee details
func GetEmployee(c *fiber.Ctx) error {

	ID := c.Params("empid")

	db := database.DB

	var employee []models.Employee = make([]models.Employee, 0)

	db.Find(c.Context(), &employee, ID)

	return c.JSON(fiber.Map{

		"status":  "success",
		"message": "Product found",
		"Emp":     employee,
	})
}
