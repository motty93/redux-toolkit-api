package handler

import (
	"app/pkg/db/model"
	"app/pkg/service"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// Login create jwt token for a login user
func (u *User) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Email or Password is empty.")
	}

	err := u.su.Session(email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Is(err, gorm.ErrRecordNotFound))
	}

	// Create jwt token
	token := jwt.New(jwt.SigningMethodES256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate encoded token.")
	}
	response := map[string]string{"token": t}

	return c.JSON(http.StatusOK, response)
}

// Restricted is jwt auth
func (u *User) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if _, ok := claims["email"]; ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unauthorized token.")
	}

	return c.String(http.StatusOK, "Welcome!")
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
