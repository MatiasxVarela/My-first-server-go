package service

import (
	"github.com/gofiber/fiber/v2"
)

/* Create new task */
func CreateUser(c *fiber.Ctx) error {
	return c.JSON("Create user")
}

/* Find all tasks */
func FindUser(c *fiber.Ctx) error {
	return c.JSON("Find all users")
}

/* Find tasks by id */
func FindOneUser(c *fiber.Ctx) error {
	return c.JSON("Find one user")
}

/* Update task */
func UpdateUser(c *fiber.Ctx) error {
	return c.JSON("Update user")
}

/* Delete Task */
func DeleteUser(c *fiber.Ctx) error {
	return c.JSON("Delete user")
}
