package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskBody struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var tasksList = make(map[int]*TaskBody)

/* Create new task */
func Create(c *fiber.Ctx) error {
	newTask := new(TaskBody)

	if err := c.BodyParser(newTask); err != nil {
		return err
	}

	newTask.Id = len(tasksList)

	tasksList[len(tasksList)] = newTask

	return c.JSON(newTask)
}

/* Find all tasks */
func Find(c *fiber.Ctx) error {
	return c.JSON(tasksList)
}

/* Find tasks by id */
func FindOne(c *fiber.Ctx) error {
	taskId := c.Params("id")

	taskIdNum, err := strconv.Atoi(taskId)

	if err != nil {
		return c.JSON("Task no found")
	}

	if len(tasksList) > taskIdNum && len(tasksList) > 0 {
		return c.JSON(tasksList[taskIdNum])
	} else {
		return c.JSON("Task no found")
	}
}
