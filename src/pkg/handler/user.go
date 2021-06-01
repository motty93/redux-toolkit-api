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

type User struct {
	su *service.User
}

func NewUserHandler(su *service.User) *User {
	return &User{su: su}
}

// GetUsers get all theirs
func (u *User) GetUsers(c echo.Context) error {
	users, err := u.su.Users()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Is(err, gorm.ErrRecordNotFound))
	}

	return c.JSON(http.StatusOK, users)
}

// GetUser find id
func (u *User) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	user, err := u.su.User(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Is(err, gorm.ErrRecordNotFound))
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser user create
func (u *User) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind user.")
	}

	err := u.su.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to insert data.")
	}

	return c.NoContent(http.StatusNoContent)
}

// UpdateUser user update
func (u *User) UpdateUser(c echo.Context) error {
	nt := new(model.User)
	if err := c.Bind(nt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind user.")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	user, err := u.su.Update(nt, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update data.")
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser user delete
func (u *User) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id parametar not found.")
	}

	err = u.su.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete data.")
	}

	return c.NoContent(http.StatusNoContent)
}
