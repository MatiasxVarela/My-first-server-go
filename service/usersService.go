package service

import (
	"database/sql"
	"myFirstServerWithGo/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/* Utils */

func findDbUser(db *sql.DB, idInt int) ([]models.Users, error) {
	var users []models.Users

	query := "SElECT * FROM users WHERE id=$1"
	rows, err := db.Query(query, idInt)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var responseUsers models.Users
		if err := rows.Scan(&responseUsers.Id, &responseUsers.FirstName, &responseUsers.LastName); err != nil {
			return users, err
		}
		users = append(users, responseUsers)
	}

	return users, nil
}

/* Create new user */
func CreateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.Users)
		if err := c.BodyParser(user); err != nil {
			return err
		}

		query := "INSERT INTO users (firstname, lastname) VALUES ($1, $2)"
		_, err := db.Exec(query, user.FirstName, user.LastName)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "Created User",
		})
	}
}

/* Find all users */
func FindAllUsers(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users []models.Users

		query := "SElECT * FROM users"

		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var responseUsers models.Users
			if err := rows.Scan(&responseUsers.Id, &responseUsers.FirstName, &responseUsers.LastName); err != nil {
				return err
			}
			users = append(users, responseUsers)
		}

		if err := rows.Err(); err != nil {
			return err
		}

		if users != nil {
			return c.JSON(users)
		} else {
			return c.Status(404).JSON(fiber.Map{
				"message": "Users not found",
			})
		}
	}
}

/* Find user by id */
func FindOneUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)

		users, err := findDbUser(db, idInt)
		if err != nil {
			return err
		}

		if users != nil {
			return c.JSON(users[0])
		} else {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found",
			})
		}
	}
}

/* Update user */
func UpdateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)

		users, err := findDbUser(db, idInt)
		if err != nil {
			return err
		}

		if len(users) < 1 {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		var usersReq models.Users
		if err := c.BodyParser(&usersReq); err != nil {
			return err
		}

		if usersReq.FirstName != "" {
			users[0].FirstName = usersReq.FirstName
		}
		if usersReq.LastName != "" {
			users[0].LastName = usersReq.LastName
		}

		return c.JSON(users)

	}
}

/* Delete user */
func DeleteUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, _ := strconv.Atoi(id)

		query := "DELETE FROM users WHERE id=$1"
		row, err := db.Exec(query, idInt)
		if err != nil {
			return err
		}

		affectedRows, _ := row.RowsAffected()

		if affectedRows > 0 {
			return c.JSON(fiber.Map{
				"message": "User Deleted",
			})
		} else {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found",
			})
		}

	}
}

/* (db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("")
	}
} */
