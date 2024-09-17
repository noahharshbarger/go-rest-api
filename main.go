package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/noahharshbarger/go-rest-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Task represents a task object
// @Description Task represents a task object
// @ID Task
// @Accept  json
// @Produce  json
type Task struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Status string `json:"status,omitempty"`
}

var tasks = []Task{
	{ID: "1", Title: "Learn Go", Status: "In Progress"},
	{ID: "2", Title: "Build a REST API", Status: "Not Started"},
}

// GetTasks retrieves all tasks
// @Summary Get all tasks
// @Description Get all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} Task
// @Router /tasks [get]
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// CreateTask adds a new task
// @Summary Create a new task
// @Description Add a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body Task true "Task"
// @Success 200 {object} Task
// @Router /tasks [post]
func CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusOK, task)
}

// @title Tasks API
// @version 1.0
// @description This is a sample API for managing tasks.
// @host localhost:8000
// @BasePath /
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/tasks", GetTasks)
	e.POST("/tasks", CreateTask)

	e.Logger.Fatal(e.Start(":8000"))
}