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

	api := e.Group("/api")
	api.GET("/tasks", th.GetTasks)
	api.GET("/tasks/:id", th.GetTask)
	api.POST("/tasks", th.CreateTask)
	api.PUT("/tasks/:id", th.UpdateTask)
	api.DELETE("/tasks/:id", th.DeleteTask)

	// us := service.NewServiceUser(db.DB)
	// uh := handler.NewUserHandler(us)
	return e
}
