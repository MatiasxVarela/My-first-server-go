package service

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
}

/* Create new user */
func CreateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return err
		}

		query := "INSERT INTO usuarios (nombre) VALUES ($1)"
		_, err := db.Exec(query, user.Name)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "Created User",
		})
	}
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
