package route

import (
	"app/pkg/db"
	"app/pkg/db/validation"
	"app/pkg/handler"
	"app/pkg/service"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Router is echo REST routing
func Router() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(handler.BodyDumper))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Validator = &validation.Custom{Validator: validator.New()}

	e.GET("/", handler.RootHandler)
	e.GET("/test", handler.TestHandler)

	ts := service.NewServiceTask(db.DB)
	th := handler.NewTaskHandler(ts)
	us := service.NewServiceUser(db.DB)
	uh := handler.NewUserHandler(us)

	e.GET("/login", uh.Login)
	a := e.Group("/api")
	a.Use(middleware.JWT([]byte("secret")))
	a.POST("", uh.Restricted)
	// GraphQL使用するのであれば以下不要
	a.GET("/tasks", th.GetTasks)
	a.GET("/tasks/:id", th.GetTask)
	a.POST("/tasks", th.CreateTask)
	a.PUT("/tasks/:id", th.UpdateTask)
	a.DELETE("/tasks/:id", th.DeleteTask)

	return e
}
