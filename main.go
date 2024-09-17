package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Status string `json:"status,omitempty"`
}

var tasks = []Task{
	{ID: "1", Title: "Learn Go", Status: "In Progress"},
	{ID: "2", Title: "Build a REST API", Status: "Not Started"},
}

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusOK, task)
}

func main() {
	e := echo.New()

	e.GET("/tasks", GetTasks)
	e.POST("/tasks", CreateTask)

	e.Logger.Fatal(e.Start(":8000"))
}