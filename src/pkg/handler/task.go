package handler

import (
	"app/pkg/db/model"
	"app/pkg/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Task struct {
	st *service.Task
}

func NewTaskHandler(st *service.Task) *Task {
	return &Task{st: st}
}

// GetTasks get all theirs
func (t *Task) GetTasks(c echo.Context) error {
	tasks, err := t.st.Tasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Is(err, gorm.ErrRecordNotFound))
	}

	return c.JSON(http.StatusOK, tasks)
}

// GetTask find id
func (t *Task) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	task, err := t.st.Task(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Is(err, gorm.ErrRecordNotFound))
	}

	return c.JSON(http.StatusOK, task)
}

// CreateTask task create
func (t *Task) CreateTask(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind task.")
	}

	err := t.st.Create(task)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to insert data.")
	}

	return c.NoContent(http.StatusNoContent)
}

// UpdateTask task update
func (t *Task) UpdateTask(c echo.Context) error {
	nt := new(model.Task)
	if err := c.Bind(nt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind task.")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	task, err := t.st.Update(nt, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update data.")
	}
	return c.JSON(http.StatusOK, task)
}

// DeleteTask task delete
func (t *Task) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	err = t.st.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete data.")
	}

	return c.NoContent(http.StatusNoContent)
}
