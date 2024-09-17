package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// setupRouter initializes a new Echo instance with the routes
func setupRouter() *echo.Echo {
	e := echo.New()
	e.GET("/tasks", GetTasks)
	e.POST("/tasks", CreateTask)
	return e
}

func TestGetTasks(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the GetTasks handler
	if assert.NoError(t, GetTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var tasks []Task
		if err := json.NewDecoder(rec.Body).Decode(&tasks); err != nil {
			t.Fatal(err)
		}

		assert.Len(t, tasks, 2) // Assuming 2 tasks in initial slice
	}
}

func TestCreateTask(t *testing.T) {
	e := setupRouter()

	task := Task{ID: "3", Title: "Test Task", Status: "Not Started"}
	body, err := json.Marshal(task)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the CreateTask handler
	if assert.NoError(t, CreateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var newTask Task
		if err := json.NewDecoder(rec.Body).Decode(&newTask); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, task, newTask)
	}
}
