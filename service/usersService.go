package service

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
}

/* Create new user */
func CreateUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*sql.DB)

	var user User
	err := c.BodyParser(&user)
	if err != nil {
	}

	// Query para insertar un usuario
	query := "INSERT INTO usuarios (nombre) VALUES ($1)"

	// Ejecutar la consulta con el nombre del usuario como parámetro
	_, err = db.Exec(query, user.Name)
	if err != nil {
		return err
	}

	return c.SendString("Usuario creado exitosamente")
}

/* Find all users */
func FindUser(c *fiber.Ctx) error {
	return c.JSON("Find all users")
}

/* Find user by id */
func FindOneUser(c *fiber.Ctx) error {
	return c.JSON("Find one user")
}

/* Update user */
func UpdateUser(c *fiber.Ctx) error {
	return c.JSON("Update user")
}

/* Delete user */
func DeleteUser(c *fiber.Ctx) error {
	return c.JSON("Delete user")
}
